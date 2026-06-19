# Projeto Korp - Desafio DevOps

Este repositório contém a solução do desafio técnico focado em Docker, Golang, Redes e Automação com Ansible. O objetivo principal é provisionar um ambiente completo de microserviços com proxy reverso e monitoramento centralizado a partir de um único comando.xz

---

## 🛠️ Tecnologias Utilizadas

* **Linguagem:** Golang 
* **Proxy Reverso:** NGINX
* **Monitoramento:** Prometheus & Grafana
* **Orquestração:** Docker & Docker Compose
* **Automação:** Ansible Playbook (IaC)

---

## 🚀 Como Executar o Projeto (Single Command)

Toda a infraestrutura foi automatizada com o Ansible. Para provisionar, buildar as imagens e subir os containers,dentro do ambiente Linux (ou WSL2) com o Ansible.

```bash
cd ansible
ansible-playbook -i inventory.ini playbook.yml

---

##  Atalho de Inicialização (Produtividade)

Caso o ambiente hospedado ou o WSL seja reiniciado, o daemon do Docker e os containers subirão desligados. Para simplificar o processo de subida da infraestrutura sem precisar digitar múltiplos comandos, criei um facilitador

### Como configurar:

1. Abra o arquivo de configuração do terminal:

```bash
sudo nano /etc/bash.bashrc
Adicione a seguinte linha ao final do arquivo (ajustando o caminho para a pasta do seu projeto):

Bash
alias subir-stack="service docker start && cd /caminho/ate/seu/projeto/ansible && ansible-playbook -i inventory.ini playbook.yml"
Salve o arquivo e recarregue as configurações para aplicar a mudança:

Bash
source /etc/bash.bashrc
Como usar no dia a dia:
Sempre que iniciar o terminal, basta alternar para o ADM e rodar o atalho:

Bash
sudo su
subir-stack
O comando iniciará o serviço do Docker, navegará automaticamente até o diretório do Ansible e reaquecerá toda a stack de forma automatizada.
