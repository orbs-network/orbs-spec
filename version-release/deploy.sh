#!/bin/bash
set -e #exit on failure

TARGET_ORG="orbsnetwork"
SOURCE_ORG="orbsnetworkstaging"
SOURCE_TAG="experimental"
TARGET_TAG=""
SOURCE_REPO="node"
TARGET_REPO=""
ROLLOUT="unspecified"
PRE_RELEASE=""


while [[ "$#" -gt 0 ]]; do
    case $1 in
        -s|--standard)       ROLLOUT="" ;;
        -h|--hotfix)         ROLLOUT="-hotfix" ;;
        -i|--immediate)      ROLLOUT="-immediate" ;;
        -c|--canary)         PRE_RELEASE="-canary" ;;
        -t|--tag)            SOURCE_TAG=$2; shift ;;
        --target-tag)        TARGET_TAG=$2; shift ;;
        -r|--repo)           SOURCE_REPO=$2; shift ;;
        --target-repo)       TARGET_REPO=$2; shift ;;
        -o|--org)            SOURCE_ORG=$2; shift ;;
        --target-org)        TARGET_ORG=$2; shift ;;
        -y)                  SKIP_CONFIRM="1" ;;
        *)
          me=`basename "$0"`
          echo "Unknown parameter passed: $1

Deployment tool for Orbs node modules. Supported module types are node services, such as signer, ethereum writer, rewards, and management node services, as well as to VChain modules.

Prerequisites:

- docker cli installed and logged in with read access to the source image, and write access to the deployment target organization and repository on the default registry
- A reference to a pre-built source docker image to deploy

The reference to the source image must be of this form:
   organization/repository:tag

If the source docker image is not present on the locally it must be available for download at the default registry

Deployment modifiers: --standard, --immediate, --hotfix and --canary.

-s, --standard Deploy randomly within a 24 hours window. This is the safest deployment mode with the highest guarantee against downtime or outages. Using --standard indicates to the Orbs node Management Service that this upgrade should be deployed in the Normal rollout window.

-i, --immediate Immediate deployment across the network. This deployment mode is unsafe for \"management-service\" and \"node\" repos. Use with caution. For VChain modules (\"node\" repo) or management-service this option is especially risky. Using --immediate indicates to the Orbs node Management Service that this upgrade should be deployed without any rollout window. Use with caution!

-h, --hotfix Deploy randomly within a 1 hour window. This compromises safety for speed of deployment and should be used only for urgent hotfixes. Using --hotfix indicates to the Orbs node Management Service that this upgrade should be deployed in the Hotfix rollout window. 

-c, --canary This option indicates deployment only to \"Canary\" VChains. For more information on Canary VChains see https://github.com/orbs-network/orbs-spec. This option is applicable only to Virtual Chain modules.

          Usage: ./$me [OPTIONS]

          -s, --standard   rolling deployment (24 hour deployment window)
          -h, --hotfix     deploy as hotfix (1 hour rolling deployment window)
          -i, --immediate  deploy immediately (no rolling deployment)
          -c, --canary     deploy only to canary vchains, relevant only for \"node\" repository images
          -t, --tag        the source tag to deploy from (default: \"experimental\")
          --target-tag     the target tag to deploy to (default: [source tag])
          -r, --repo       the source repository to deploy from (default: \"node\")
          --target-repo    the target repository to deploy to (default: [source repository])
          -o, --org        the source organization to deploy from (default: orbsnetworkstaging)
          --target-org     the target organization to deplot to (default: orbsnetwork)

          -y               suppress confirmations


          "
          exit 1 ;;
    esac
    shift
done

# ensure rolling strategy is specified when confirmation is suppressed
if [[ "$SKIP_CONFIRM" == "1" && "$ROLLOUT" == "unspecified" ]]
then
	echo "when confirmation is suppressed a deployment strategy must be specified, aborting"
	exit 2
fi

# prompt for rolling deployment policy
if [[ "$ROLLOUT" == "unspecified" ]]
then
   echo 
   echo "deployment policy unspecified. please specify a deployment policy:"
   echo
   echo "  'i'     - immediate: no rolling deployment. CAUTION!!
            Never use this policy for VChain modules (\"node\" repo), or for Management Service (\"management-service\" repo) 
            if it includes a boyar binary version upgrade or a VChain configuration change. 
            In case of doubt its recommended to avoid this policy"
   echo
   echo "  'h'     - hotfix:    1 hour depoyment window"
   echo "  any key - standard:  24 hour deployment window"
   echo    
   read -p "Enter your selection in a single character: " -n 1 -r
   echo    # (optional) move to a new line

   if [[ $REPLY =~ ^[iI]$ ]]
   then 
      echo "Immediate deployment selected"
      ROLLOUT="-immediate"
   elif [[ $REPLY =~ ^[hH]$ ]]
   then
      echo "Hotfix deployment selected"
      ROLLOUT="-hotfix"
   else 
      echo "Standard deployment selected"
      ROLLOUT=""
   fi
fi

# target tag defaults to source tag, target repository defaults to source repository
if [ "$TARGET_TAG" = "" ]; then  TARGET_TAG=$SOURCE_TAG ; fi
if [ "$TARGET_REPO" = "" ]; then TARGET_REPO=$SOURCE_REPO ; fi

SOURCE_FULL="$SOURCE_ORG/$SOURCE_REPO:$SOURCE_TAG"
TARGET_FULL="$TARGET_ORG/$TARGET_REPO:$TARGET_TAG$PRE_RELEASE$ROLLOUT"

echo "Deploying: $SOURCE_FULL --> $TARGET_FULL"
[ "$PRE_RELEASE" != "" ] && echo "Selected deployment group $PRE_RELEASE"
[ "$ROLLOUT"     != "" ] && echo "Selected rollout window indicator $ROLLOUT"

echo
echo
[ "$SOURCE_REPO" != "$TARGET_REPO" ] && echo "Warning!!! The source and target repository names ($SOURCE_REPO, $TARGET_REPO) do not match. Was it intended?"

if [ "$SKIP_CONFIRM" != "1" ]
then
   read -p "Proceed with deployment of target image ($TARGET_FULL)? [y/N] " -n 1 -r
   echo    # (optional) move to a new line

   if [[ ! $REPLY =~ ^[Yy]$ ]]
   then
      echo "aborting..."
      exit 1
   fi
fi

docker pull $SOURCE_FULL
docker tag $SOURCE_FULL $TARGET_FULL
docker push $TARGET_FULL

echo
echo "Successfully deployed $TARGET_FULL"
