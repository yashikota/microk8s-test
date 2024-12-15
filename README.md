# microk8s-test

## Ansible

```sh
pipx install --include-deps ansible
cd ansible
ansible-playbook -i inventory.yaml playbook.yaml -e "remote_user={username}"
```
