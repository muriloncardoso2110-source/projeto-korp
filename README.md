# Projeto Korp - Desafio DevOps

Este repositório contém a solução do desafio técnico focado em Docker, Golang, Redes e Automação com Ansible. O objetivo principal é provisionar um ambiente completo de microserviços com proxy reverso e monitoramento centralizado a partir de um único comando.

---

## 🛠️ Tecnologias Utilizadas

* **Linguagem:** Golang 
* **Proxy Reverso:** NGINX
* **Monitoramento:** Prometheus & Grafana
* **Orquestração:** Docker & Docker Compose
* **Automação:** Ansible Playbook (IaC)

---

## 🚀 Como Executar o Projeto (Single Command)

Toda a infraestrutura foi automatizada com o Ansible. Para provisionar, buildar as imagens e subir os containers,dentro do ambiente Linux (ou WSL2) com o Ansible instalado e é só executar:

```bash
cd ansible
ansible-playbook -i inventory.ini playbook.yml
