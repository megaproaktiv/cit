pushd .
cd /tmp
rm -f awscliv2.zip
rm -rf ./aws
curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip"
unzip awscliv2.zip
sudo ./aws/install
popd 
brew install go-task/tap/go-task
