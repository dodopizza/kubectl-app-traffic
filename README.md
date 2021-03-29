# kubectl app-traffic

Kubectl plugin, that enables / disables traffic to kubernetes pods.
It's helpful during incident or application outage

## Installation

Using curl:

```bash
curl -LO https://raw.githubusercontent.com/dodopizza/kubectl-app-traffic/main/bin/kubectl-app_traffic
chmod +x kubectl-app_traffic
mv ./kubectl-app_traffic /usr/local/bin/kubectl-app_traffic
```

## Usage

```bash
# Generic invocation
kubectl app-traffic -n <namespace> <--enable|--disable> <service>

# Disable traffic from service 'foo' located in namespace 'bar'
kubectl app-traffic -n bar --disable foo

# Enable traffic to service 'foo' (after it was disabled) located in namespace 'bar'
kubectl app-traffic -n bar --enable foo

# Namespace flag can be omitted, in this case it will be used from current kube config
# Disable traffic from service `foo`
kubectl app-traffic --enable foo
```
