
REMINDER: AUTOBUILD ENABLED.

**GKE ARCHITECTURE**
>Master
 * (API Server: Controller Manager | Scheduler) --- (etcd)
 * NODE | NODE | ...
>NODE
* Kubelet
* Kube Proxy :<<:<-- USER POC
* {Pods}

----------------------------------------------------------------------------
==========SIMPLE START=======
----------------------------------------------------------------------------
via CLOUD SHELL
-----------------
1. (CONFIGURE PROJECT/GIT) 
    * REMINDER TO-DO: If x1adoc-ml/almightyai/systemctl hive is not enabled, do alt implementation.

2. Set Identifiers
    * `gcloud config set project #`
    * `gcloud config set compute/zone us-central1`
    * `gcloud config set compute/zone us-central1-a`

3. Enable APIs
    * `gcloud services enable container.googleapis.com`
    * `gcloud services enable containerregistry.googleapis.com`

4. CLONE GIT REPO > WORKINDIR

DOCKER PACKAGING (V1) & CONTAINER REGISTRY UPLINK
---------------------------------------------------
* `docker build -t gcr.io/A/B:v1 .`
* `gcloud auth configure-docker` :auth gcloud as cred helper
* `docker push gcr.io/A/B:v1` 

GKE CLUSTER NODE GENERATION (CLUSTERPHOBIC == a pool of node pools which are essentially a set of MIGs)
-------------------------------------------
*  NOTE: Node Image not to be confused with container image (pod scope).
* `gcloud create deployment CLUSTERPHOBIC --image=gcr.io/A/B:v1`
*  kubectl auth ---> CLUSTERPHOBIC

17:00 Zulu Time
--------------------
* `kubectl run CLUSTERPHOBIC --image=gcr.io/A/B:v1 --port 80`
* `kubectl get pods`
*  Load Balance w/ Ingress >> Expose IP 

**ASSUMING: Network Load Balance (TCP XOR UDP)**
---------------------------------------------
>(1.Seattle 2.NYC 3.London)
-  When a user in London connects into the U.S. West backend, the traffic ingresses closest to London, because the range is anycasted. 
- You need to forward the original packets unproxied && response is directly sent back to user. (!= ingress route)
- NLB uses target pools for session affinity (not backend).

----------------------------------------------------------------------------
========== PREVIOUS DETAILED ROUTE =======
----------------------------------------------------------------------------

Overview
---------
BUILD IT
1. Package app > Docker Image
2. Run container locally (optional)
STORE IT
3. Upload image to registry
4. Create Container Cluster
5. Deploy App to Cluster
6. Expose App to Public
7. Scale Up Deployment
8. Deploy New Version of App


PreReq:
----------
- GCP Project
- Enable Billing (note: set alerts)
- Enable Kubernetes Engine API (APIs & Services > Enable > Search Kubernetes)
- GCloud SDK 
- `gcloud components install kubectl`
- Docker CE:  `docker -v`
- Git


BUILD IT
---------------
- `export PROJECT_ID=[PROJECT_ID]` 
- `docker build --tag=gcr.io/#PROJECT_ID#/#APP_FOLDER#:v1 .`
    - gcr.io = Google Cloud Registry
    - verify build `docker images`

UPLOAD (Container Registry)
--------------
- First Time Users: `gcloud auth configure-docker`
- upload by running: `docker push gcr.io/#PROJECT_ID#/#APP_FOLDER#:v1`

LOCAL RUN:
--------------
- `docker run --rm -p 8080:8080 gcr.io/#PROJECT_ID#/#APP_FOLDER#:v1`
- `curl localhost:800`


CREATE CLUSTER
-----------------
- Cluster=Pool of Compute Engine VM instances (running Kubernetes)

- `gcloud config set project #PROJECT_ID#`
- `gcloud config set compute/zone [COMPUTE_ENGINE_ZONE]`
    - To view list: `gcloud compute zones list`

- `gcloud container clusters create #CLUSTERNAME# --num-nodes=2`
    - verify: `gcloud compute instances list`


DEPLOY APPLICATION
-----------------
- DEPLOY `kubectl create deployment #CLUSTERNAME#--image=gcr.io/#PROJECT_ID#/#APP_FOLDER#:v1`
- SEE PODS `kubectl get pods`


LOAD BALANCE & EXPOSE TO PUBLIC (subject to billing)
-------------------
- `kubectl expose deployment #CLUSTERNAME# --type=LoadBalancer --port 80 --target-port 8080`


SCALE_UP
----------
- `kubectl scale deployment hello-web --replicas=3`
- Verify: `kubectl get deployment hello-web`
- Verify: `kubectl get pods`



NEW VERSION DEPLYMENT
-----------------------
- `docker build --tag=gcr.io/#PROJECT_ID#/#APP_FOLDER#:v2 .`
- `docker push gcr.io/#PROJECT_ID#/#APP_FOLDER#:v2`
- `kubectl set image deployment/#CLUSTERNAME# #APP_FOLDER#=gcr.io/#PROJECT_ID#/#APP_FOLDER#:v2`

- VERIFY: http://[EXTERNAL_IP]


CLEANUP
----------
- rm service: `kubectl delete service hello-web`
- rm cluster: `gcloud container clusters delete hello-cluster`


TODO: 
---------------
1. APIs & Services: Create Credentials for Kuberenetes API




