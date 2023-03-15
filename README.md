# kubectl app-traffic

Kubectl plugin, that enables / disables traffic to kubernetes pods.
It's helpful during incident or application outage

## Installation

### Using curl

OS X:

```bash
curl -LO https://github.com/dodopizza/kubectl-app-traffic/releases/latest/download/kubectl-app_traffic_Darwin_x86_64.tar.gz
tar -xvf kubectl-app_traffic_Darwin_x86_64.tar.gz
chmod +x kubectl-app_traffic
mv ./kubectl-app_traffic /usr/local/bin/kubectl-app_traffic
```

Linux:

```bash
curl -LO https://github.com/dodopizza/kubectl-app-traffic/releases/latest/download/kubectl-app_traffic_Linux_x86_64.tar.gz
tar -xvf kubectl-app_traffic_Linux_x86_64.tar.gz
chmod +x kubectl-app_traffic
mv ./kubectl-app_traffic /usr/local/bin/kubectl-app_traffic
```

Windows:

```powershell
Invoke-WebRequest -Uri "https://github.com/dodopizza/kubectl-app-traffic/releases/latest/download/kubectl-app_traffic_Windows_x86_64.zip" -OutFile "kubectl-app_traffic_Windows_x86_64.zip"
Expand-Archive -Path "kubectl-app_traffic_Windows_x86_64.zip"
Move-Item -Path "kubectl-app_traffic_Windows_x86_64/kubectl-app_traffic_Windows_x86_64.exe" -Destination "$env:USERPROFILE/.kube/plugins"
```

## Usage

```bash
# Generic invocation
kubectl app-traffic -n <namespace> <enable|disable> <service|ingress> <service_name|ingress_name>

# Disable traffic from service 'foo' located in namespace 'bar'
kubectl app-traffic -n bar disable service foo

# Enable traffic to service 'foo' (after it was disabled) located in namespace 'bar'
kubectl app-traffic -n bar enable service foo

# Disable traffic from ingress 'foo' located in namespace 'bar'
kubectl app-traffic -n bar disable ingress foo

# Enable traffic to ingress 'foo' (after it was disabled) located in namespace 'bar'
kubectl app-traffic -n bar enable ingress foo

# Namespace flag can be omitted, in this case it will be used from current kube config
# Disable traffic from ingress `foo`
kubectl app-traffic disable foo
```
