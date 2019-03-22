# -*- mode: ruby -*-
# vi: set ft=ruby :

# All Vagrant configuration is done below. The "2" in Vagrant.configure
# configures the configuration version (we support older styles for
# backwards compatibility). Please don't change it unless you know what
# you're doing.


$k8s_installation_script = <<-SCRIPT
export LC_ALL="en_US.UTF-8"
export LC_CTYPE="en_US.UTF-8"
curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key add -
echo 'deb http://apt.kubernetes.io/ kubernetes-xenial main' | sudo tee /etc/apt/sources.list.d/kubernetes.list
sudo apt-get update -y
sudo apt-get install kubelet kubeadm kubectl -y
SCRIPT

$k8s_cluster_join_script = <<-SCRIPT
export LC_ALL="en_US.UTF-8"
export LC_CTYPE="en_US.UTF-8"
echo "Bootstrap script running......"
SCRIPT

Vagrant.configure("2") do |config|
  config.vm.provision "docker"
  config.vm.provision "shell", inline: $k8s_installation_script

  config.vm.define "master" do |master|
      master.vm.box = "bento/ubuntu-16.04"
      master.vm.hostname = "master"
      master.vm.network "private_network", ip: "192.168.33.10"
  end

  config.vm.define "node_1" do |node_1|
    node_1.vm.box = "bento/ubuntu-16.04"
    node_1.vm.hostname = "node-1"
    node_1.vm.network "private_network", ip: "192.168.33.11"
    # node_1.vm.provision "shell", inline: $k8s_cluster_join_script
  end

end