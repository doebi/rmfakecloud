export DOCKER_BUILDKIT=1
docker build --network host -t doebi/rmfakecloud --no-cache .
