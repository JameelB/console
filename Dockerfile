FROM busybox
ADD frontend/public/dist /opt/bridge/assets
ADD bin/bridge /opt/bridge/bin/bridge
ADD schema/v1.json /opt/bridge/v1.json
CMD /opt/bridge/bin/bridge --public-dir=/opt/bridge/assets --address=0.0.0.0:9090 --k8s-url=http://localhost:8080 --discovery-json=/opt/bridge/v1.json