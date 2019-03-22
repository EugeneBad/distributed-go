# -*- mode: ruby -*-
# vi: set ft=ruby :

# All Vagrant configuration is done below. The "2" in Vagrant.configure
# configures the configuration version (we support older styles for
# backwards compatibility). Please don't change it unless you know what
# you're doing.
Vagrant.configure("2") do |config|
  config.vm.define "node_1" do |node_1|
    node_1.vm.box = "bento/ubuntu-16.04"
    node_1.vm.hostname = "node-1"
    node_1.vm.network "private_network", ip: "192.168.33.10"
  end
  config.vm.provision "docker"
end