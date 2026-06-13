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

Evidências do Funcionamento
Aqui estão as capturas de tela do ambiente totalmente operacional:

<img width="1867" height="1002" alt="image" src="https://github.com/user-attachments/assets/fdc9fa97-1ae3-4c5c-9c46-4025c4620f1f" />

<img width="1874" height="1008" alt="image" src="https://github.com/user-attachments/assets/02d28efa-fbd8-454f-898c-9d821a6b7684" />

<img width="1868" height="1010" alt="image" src="https://github.com/user-attachments/assets/f30ef9aa-72bd-4f14-95a6-bfd3ea74c862" />

 API respondendo o JSON com horário UTC dinâmico (Porta 8081)
 Alvos (Targets) ativos e coletando no Prometheus (Porta 9090)
 Dashboard do Grafana monitorando os recursos em tempo real (Porta 3000)


