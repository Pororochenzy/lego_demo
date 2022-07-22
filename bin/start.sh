./stup.sh start console console ./conf/console.yaml
sleep 1
./stup.sh start gate gate ./conf/gate.yaml 
sleep 1
./stup.sh start live live ./conf/live.yaml 
