```bash
docker build --platform linux/amd64 -t secret104278/cfs-test .
docker push secret104278/cfs-test
kubectl delete -f job.yaml; kubectl apply -f job.yaml
kubectl get pods -l job-name=cfs-test
kubectl logs -f -l job-name=cfs-test
```
