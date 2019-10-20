Packaging & Deploying Containerized Web App on GKE
TRIGGERS:
------------
- Docker Autobuild triggers a new build with every git push to this repository.

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




