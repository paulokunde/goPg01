-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema infra
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema infra
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `infra` DEFAULT CHARACTER SET utf8 ;
USE `infra` ;

-- -----------------------------------------------------
-- Table `infra`.`pregao`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `infra`.`pregao` ;

CREATE TABLE IF NOT EXISTS `infra`.`pregao` (
  `username` VARCHAR(16) NULL,
  `create_time` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` TIMESTAMP NULL,
  `numero` INT NOT NULL,
  `ano` INT NOT NULL,
  PRIMARY KEY (`numero`, `ano`));


-- -----------------------------------------------------
-- Table `infra`.`solicitante`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `infra`.`solicitante` ;

CREATE TABLE IF NOT EXISTS `infra`.`solicitante` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `nome` VARCHAR(45) NOT NULL,
  `create_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` TIMESTAMP NULL,
  `usuario` VARCHAR(16) NULL,
  PRIMARY KEY (`id`));


-- -----------------------------------------------------
-- Table `infra`.`tipo`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `infra`.`tipo` ;

CREATE TABLE IF NOT EXISTS `infra`.`tipo` (
  `username` VARCHAR(16) NULL,
  `create_time` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  `id` INT UNIQUE NOT NULL AUTO_INCREMENT,
  `nome` VARCHAR(128) NULL,
  PRIMARY KEY (`id`));


-- -----------------------------------------------------
-- Table `infra`.`marca`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `infra`.`marca` (
  `username` VARCHAR(16) NOT NULL,
  `create_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `id` INT UNIQUE NOT NULL AUTO_INCREMENT,
  `nome` VARCHAR(128) NOT NULL,
  PRIMARY KEY (`id`));


-- -----------------------------------------------------
-- Table `infra`.`modelo`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `infra`.`modelo` ;

CREATE TABLE IF NOT EXISTS `infra`.`modelo` (
  `username` VARCHAR(16) NULL,
  `create_time` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  `id` INT UNIQUE NOT NULL AUTO_INCREMENT,
  `nome` VARCHAR(128) NULL,
  `marca_id` INT UNIQUE NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_modelo_marca1_idx` (`marca_id` ASC) VISIBLE,
  CONSTRAINT `fk_modelo_marca1`
    FOREIGN KEY (`marca_id`)
    REFERENCES `infra`.`marca` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION);


-- -----------------------------------------------------
-- Table `infra`.`equipamento`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `infra`.`equipamento` ;

CREATE TABLE IF NOT EXISTS `infra`.`equipamento` (
  `username` VARCHAR(16) NULL,
  `create_time` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  `id` INT UNIQUE NOT NULL AUTO_INCREMENT,
  `nome` VARCHAR(128) NULL,
  `tipo_id` INT UNIQUE NOT NULL,
  `modelo_id` INT UNIQUE NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_equipamento_tipo1_idx` (`tipo_id` ASC) VISIBLE,
  INDEX `fk_equipamento_modelo1_idx` (`modelo_id` ASC) VISIBLE,
  CONSTRAINT `fk_equipamento_tipo1`
    FOREIGN KEY (`tipo_id`)
    REFERENCES `infra`.`tipo` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_equipamento_modelo1`
    FOREIGN KEY (`modelo_id`)
    REFERENCES `infra`.`modelo` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION);


-- -----------------------------------------------------
-- Table `infra`.`componente`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `infra`.`componente` ;

CREATE TABLE IF NOT EXISTS `infra`.`componente` (
  `username` VARCHAR(16) NULL,
  `create_time` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  `id` INT UNIQUE NOT NULL AUTO_INCREMENT,
  `nome` VARCHAR(128) NULL,
  `modelo_id` INT UNIQUE NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_componente_modelo1_idx` (`modelo_id` ASC) VISIBLE,
  CONSTRAINT `fk_componente_modelo1`
    FOREIGN KEY (`modelo_id`)
    REFERENCES `infra`.`modelo` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION);


-- -----------------------------------------------------
-- Table `infra`.`pecas`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `infra`.`pecas` ;

CREATE TABLE IF NOT EXISTS `infra`.`pecas` (
  `equipamento_id` INT UNIQUE NOT NULL,
  `componente_id` INT UNIQUE NOT NULL,
  `create_time` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  `username` VARCHAR(45) NULL,
  INDEX `fk_pecas_equipamento1_idx` (`equipamento_id` ASC) VISIBLE,
  INDEX `fk_pecas_componente1_idx` (`componente_id` ASC) VISIBLE,
  CONSTRAINT `fk_pecas_equipamento1`
    FOREIGN KEY (`equipamento_id`)
    REFERENCES `infra`.`equipamento` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_pecas_componente1`
    FOREIGN KEY (`componente_id`)
    REFERENCES `infra`.`componente` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION);


-- -----------------------------------------------------
-- Table `infra`.`itens`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `infra`.`itens` ;

CREATE TABLE IF NOT EXISTS `infra`.`itens` (
  `data` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `pregao_numero` INT NOT NULL,
  `pregao_ano` INT NOT NULL,
  `equipamento_id` INT UNIQUE NOT NULL,
  `solicitante_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_itens_pregao1_idx` (`pregao_numero` ASC, `pregao_ano` ASC) VISIBLE,
  INDEX `fk_itens_equipamento1_idx` (`equipamento_id` ASC) VISIBLE,
  INDEX `fk_itens_solicitante1_idx` (`solicitante_id` ASC) VISIBLE,
  CONSTRAINT `fk_itens_pregao1`
    FOREIGN KEY (`pregao_numero` , `pregao_ano`)
    REFERENCES `infra`.`pregao` (`numero` , `ano`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_itens_equipamento1`
    FOREIGN KEY (`equipamento_id`)
    REFERENCES `infra`.`equipamento` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_itens_solicitante1`
    FOREIGN KEY (`solicitante_id`)
    REFERENCES `infra`.`solicitante` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION);


-- -----------------------------------------------------
-- Table `infra`.`fornecedor`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `infra`.`fornecedor` ;

CREATE TABLE IF NOT EXISTS `infra`.`fornecedor` (
  `username` VARCHAR(16) NULL,
  `create_time` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  `id` INT UNIQUE NOT NULL AUTO_INCREMENT,
  `nome` VARCHAR(128) NULL,
  PRIMARY KEY (`id`));


-- -----------------------------------------------------
-- Table `infra`.`analise`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `infra`.`analise` ;

CREATE TABLE IF NOT EXISTS `infra`.`analise` (
  `atende` TINYINT NULL,
  `obs` TEXT(256) NULL,
  `create_time` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  `username` VARCHAR(45) NULL,
  `id` INT NOT NULL AUTO_INCREMENT,
  `itens_id` INT UNSIGNED NOT NULL,
  `fornecedor_id` INT UNIQUE NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_analise_itens1_idx` (`itens_id` ASC) VISIBLE,
  INDEX `fk_analise_fornecedor1_idx` (`fornecedor_id` ASC) VISIBLE,
  CONSTRAINT `fk_analise_itens1`
    FOREIGN KEY (`itens_id`)
    REFERENCES `infra`.`itens` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_analise_fornecedor1`
    FOREIGN KEY (`fornecedor_id`)
    REFERENCES `infra`.`fornecedor` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION);


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
