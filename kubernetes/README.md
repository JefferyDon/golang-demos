# Kubernetes Demos

This directory stores all demo program related to kubernetes, most of the programs are the usage of client-go SDK.

# 1. Prerequisite

Before you running programs in here, make sure `${HOME}/.kube/config` file exists, 
and make sure this kube-config can access the exactly kubernetes cluster you are using for test.

> All kubernetes client created in programs are using this file. Based on Kubernetes v1.23.9, using client-go SDK v0.23.9 
