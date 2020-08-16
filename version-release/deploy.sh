#!/bin/bash
set -e #exit on failure

TARGET_ORG="orbsnetwork"
SOURCE_ORG="orbsnetworkstaging"
SOURCE_TAG="experimental"
TARGET_TAG=""
SOURCE_REPO="node"
TARGET_REPO=""
ROLLOUT="normal"
PRE_RELEASE=""


while [[ "$#" -gt 0 ]]; do
    case $1 in               
        -h|--hotfix)         ROLLOUT="hotfix" ;;
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

VChain modules:
For VChain modules two options modify their deployment: --hotfix or --canary. 

-h, --hotfix To ensure constant availability of a quorum of nodes, the Orbs Management Service deploys updates to VChain core modules with a gradual rollout window. Orbs supports two mode of rollout: Normal for a safer and longer rollout window of 24 hours (by default), and Hotfix rollout window for urgent, expedited, rollouts within 1 hour (by default). for more information see https://github.com/orbs-network/orbs-spec 
Using --hotfix indicates to the Orbs node Management Service that this upgrade should be deployed in the Hotfix rollout window. it is only applicable to Virtual Chain modules.

-c, --canary This option indicates deployment only to "Canary" VChains. For more information on Canary VChains see https://github.com/orbs-network/orbs-spec. This option is applicable only to Virtual Chain modules. Using this option will result in a "canary" pre-release indicator appended to the semver tag of the target image.

          Usage: ./$me [OPTIONS] 
          
          -h, --hotfix     deploy as hotfix (quick deployment), relevant only for "node" repository images
          -c, --canary     deploy only to canary vchains, relevant only for "node" repository images
          -t, --tag        the source tag to deploy from (default: "experimental")
          --target-tag     the target tag to deploy to (default: [source tag])
          -r, --repo       the source repository to deploy from (default: "node")
          --target-repo    the target repository to deploy to (default: [source repository])
          -o, --org        the source organization to deploy from (default: orbsnetworkstaging)
          --target-org     the target organization to deplot to (default: orbsnetwork)
          
          -y               suppress confirmation

          Use this tool to publish docker images to docker repository.
          
          "
          exit 1 ;;
    esac
    shift
done

# target tag defaults to source tag, target repository defaults to source repository
if [ "$TARGET_TAG" = "" ]; then  TARGET_TAG=$SOURCE_TAG ; fi
if [ "$TARGET_REPO" = "" ]; then TARGET_REPO=$SOURCE_REPO ; fi
 
SOURCE_FULL="$SOURCE_ORG/$SOURCE_REPO:$SOURCE_TAG"
TARGET_FULL="$TARGET_ORG/$TARGET_REPO:$TARGET_TAG$PRE_RELEASE"

echo "Deploying: $SOURCE_FULL --> $TARGET_FULL"
echo "Rollout as $ROLLOUT"
echo

if [ "$SKIP_CONFIRM" != "1" ]
then
   echo 
   REPLY=""
   read -p "Proceed with building the target image ($TARGET_FULL)? " -n 1 -r
   echo    # (optional) move to a new line 

   if [[ ! $REPLY =~ ^[Yy]$ ]]
   then
      echo "aborting..."
      exit 1
   fi
fi

VER_L="network.orbs.deploy-tool-version=1.0"
ROL_L="network.orbs.rollout-mode=$ROLLOUT"
DOCKERFILE="FROM $SOURCE_ORG/$SOURCE_REPO:$SOURCE_TAG"

# create and tag the new docker image for deployment
echo $DOCKERFILE | docker build --pull --label $VER_L --label $ROL_L --tag $TARGET_ORG/$TARGET_REPO:$TARGET_TAG$PRE_RELEASE -


if [ "$SKIP_CONFIRM" != "1" ]
then
   echo
   REPLY=""
   read -p "Proceed with pushing the target image ($TARGET_FULL)? " -n 1 -r
   echo    # (optional) move to a new line

   if [[ ! $REPLY =~ ^[Yy]$ ]]
   then
      echo "aborting..."
      exit 1
   fi
fi

# upload the image to target organization, repository and tag
docker push $TARGET_ORG/$TARGET_REPO:$TARGET_TAG$PRE_RELEASE

