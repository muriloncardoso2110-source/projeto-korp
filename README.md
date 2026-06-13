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

 API respondendo o JSON com horário UTC dinâmico (Porta 8081)
 <img width="1880" height="1035" alt="image3" src="https://github.com/user-attachments/assets/3e899f39-aa17-4c21-966b-6be6c5f80d7a" />

 Alvos (Targets) ativos e coletando no Prometheus (Porta 9090)
 <img width="1875" height="1002" alt="Image1" src="https://github.com/user-attachments/assets/e87c2ec6-7eeb-4287-9a2c-b6250dce2db9" />
<img width="1877" height="1007" alt="image2" src="https://github.com/user-attachments/assets/3d7bd128-cc61-4883-99ee-362ad3bc80a0" />

 Dashboard do Grafana monitorando os recursos em tempo real (Porta 3000)
![Uploading Image1.png…]()


