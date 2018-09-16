python3 -m pip install .
echo '***********************************************'
echo 'setting up heart-risk'
mkdir -p /opt/heart-risk/
echo 'directory made /opt/heart-risk'
cp -R lib/* /opt/heart-risk/
echo 'files copied to /opt/heart-risk/'
CELERY_DIR="$(which celery)"
echo $CELERY_DIR
echo "[Unit]
Description=CELERY WORKERS
After=network.service
[Service]
WorkingDirectory=/opt/heart-risk/
User=nobody
ExecStart=$CELERY_DIR -A task worker --loglevel=info
[Install]
WantedBy=multi-user.target" >> 'heart-risk.service'
cp heart-risk.service /etc/systemd/system/heart-risk.service
echo 'Removing temp service file.....'
#rm heart-risk.service
systemctl enable heart-risk.service
service heart-risk.service start
service heart-risk.service status -l
echo 'Installed and started heart-risk'
echo '***********************************************'