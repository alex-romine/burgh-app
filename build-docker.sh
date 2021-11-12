set -e

IMAGE_NAME="$1"
IMAGE_VERSION="0.1.$(date +%Y%m%d%I%M%S)"
GOOS="$2"
GOARCH="$3"

if [ -z "$IMAGE_NAME" ]
then
  echo "need image name"
  exit 1
fi

echo "building image"
docker build . -t "aromine2/$IMAGE_NAME:$IMAGE_VERSION"

echo "pushing image"
docker push "aromine2/$IMAGE_NAME:$IMAGE_VERSION"
