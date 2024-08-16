terraform {
  backend "local" {
    path = "./backend/terraform.tfstate"
  }
}

provider "aws" {
  region = "us-east-2"
}

resource "aws_instance" "web-server-instance" {
  ami           = "ami-0862be96e41dcbf74"
  instance_type = "t2.micro"

  tags = {
    Name = "ani-test-web-server"
    project = "awi-infra-guard"
    ci_created = "true"
  }
}

resource "aws_vpc" "test-vpc" {
  cidr_block = "10.0.0.0/16"
  tags = {
    Name = "ani-test-vpc"
    project = "awi-infra-guard"
    ci_created = "true"
  }
}

resource "aws_internet_gateway" "test-gw" {
  vpc_id = aws_vpc.test-vpc.id
  tags = {
    Name = "ani-test-internet-gateway"
    project = "awi-infra-guard"
    ci_created = "true"
  }
}

resource "aws_route_table" "test-route-table" {
  vpc_id = aws_vpc.test-vpc.id

  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.test-gw.id
  }

  route {
    ipv6_cidr_block = "::/0"
    gateway_id      = aws_internet_gateway.test-gw.id
  }

  tags = {
    Name = "ani-test-route-table"
    project = "awi-infra-guard"
    ci_created = "true"
  }
}

resource "aws_subnet" "test-subnet" {
  vpc_id            = aws_vpc.test-vpc.id
  cidr_block        = "10.0.1.0/24"
  availability_zone = "us-east-2a"

  tags = {
    Name = "ani-test-subnet"
    project = "awi-infra-guard"
    ci_created = "true"
  }
}