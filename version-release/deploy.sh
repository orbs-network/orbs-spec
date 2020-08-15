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
          echo -e "Unknown parameter passed: $1 \n\
          \n\
          Usage: ./$me [OPTIONS] \n\
          \n\
          -h, --hotfix     deploy as hotfix (quick deployment), relevant only for "node" repository images \n\
          -c, --canary     deploy only to canary vchains, relevant only for "node" repository images \n\
          -t, --tag        the source tag to deploy from \(default: "experimental"\) \n\
          --target-tag     the target tag to deploy to \(default: [source tag]\) \n\
          -r, --repo       the source repository to deploy from \(default: "node"\) \n\
          --target-repo    the target repository to deploy to \(default: [source repository]\) \n\
          -o, --org        the source organization to deploy from \(default: orbsnetworkstaging\) \n\
          --target-org     the target organization to deplot to \(default: orbsnetwork\) \n\
          \n\
          -y               suppress confirmation \n\
          "
          exit 1 ;;
    esac
    shift
done

# target tag defaults to source tag, target repository defaults to source repository
if [ "$TARGET_TAG" = "" ]; then  TARGET_TAG=$SOURCE_TAG ; fi
if [ "$TARGET_REPO" = "" ]; then TARGET_REPO=$SOURCE_REPO ; fi
 

echo "Deploying: $SOURCE_ORG/$SOURCE_REPO:$SOURCE_TAG --> $TARGET_ORG/$TARGET_REPO:$TARGET_TAG$PRE_RELEASE"
echo "Rollout as $ROLLOUT"
echo

if [ "$SKIP_CONFIRM" != "1" ]
then
   read -p "Are you sure?" -n 1 -r
   echo    # (optional) move to a new line 

   if [[ ! $REPLY =~ ^[Yy]$ ]]
   then
      echo "cancelling..."
      exit 1
   fi
fi

VER_L="network.orbs.deploy-tool-version=1.0"
ROL_L="network.orbs.rollout-mode=$ROLLOUT"
DOCKERFILE="FROM $SOURCE_ORG/$SOURCE_REPO:$SOURCE_TAG"

# create and tag the new docker image for deployment
echo $DOCKERFILE | docker build --pull --label $VER_L --label $ROL_L --tag $TARGET_ORG/$TARGET_REPO:$TARGET_TAG$PRE_RELEASE -

# upload the image to target organization, repository and tag
docker push $TARGET_ORG/$TARGET_REPO:$TARGET_TAG$PRE_RELEASE

