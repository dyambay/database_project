-- MySQL Script generated by MySQL Workbench
-- 04/19/15 21:45:37
-- Model: New Model    Version: 1.0
-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='TRADITIONAL,ALLOW_INVALID_DATES';

-- -----------------------------------------------------
-- Schema PathfinderEncounter
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema PathfinderEncounter
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `PathfinderEncounter` DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci ;
USE `PathfinderEncounter` ;

-- -----------------------------------------------------
-- Table `PathfinderEncounter`.`Monster`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `PathfinderEncounter`.`Monster` (
  `idMonster` INT NOT NULL,
  `Name` VARCHAR(45) NOT NULL,
  `CR` INT NOT NULL,
  `Alignment` VARCHAR(20) NOT NULL,
  `Size` INT NOT NULL,
  `Class` VARCHAR(25) NULL,
  `TypeName` VARCHAR(45) NOT NULL,
  `Initiative` INT NOT NULL,
  `Armor` INT NULL,
  `Shield` INT NULL,
  `Deflection` INT NULL,
  `SizeAC` INT NULL,
  `NaturalArmor` INT NULL,
  `Dodge` INT NULL,
  `MiscAC` INT NULL,
  `HitDie` INT NOT NULL,
  `Fort` INT NOT NULL,
  `Reflex` INT NOT NULL,
  `Will` INT NOT NULL,
  `BaseSpeed` INT NOT NULL,
  `Space` INT NOT NULL,
  `Reach` INT NULL,
  `Spell-Like Abilities` VARCHAR(1500) NULL,
  `Spells` VARCHAR(1500) NULL,
  `Str` INT NULL,
  `Dex` INT NULL,
  `Con` INT NULL,
  `Inte` INT NULL,
  `Wis` INT NULL,
  `Cha` INT NULL,
  `BaseAttack` INT NOT NULL,
  `CMB` VARCHAR(45) NULL,
  `CMD` VARCHAR(45) NULL,
  `Feats` VARCHAR(1500) NULL,
  `Skills` VARCHAR(1500) NULL,
  `Languages` VARCHAR(200) NULL,
  `Special Attacks` VARCHAR(1500) NULL,
  `Environment` VARCHAR(45) NOT NULL,
  `Attack1` VARCHAR(45) NULL,
  `Attack2` VARCHAR(45) NULL,
  `Attack3` VARCHAR(45) NULL,
  `Attack4` VARCHAR(45) NULL,
  `Attack5` VARCHAR(45) NULL,
  PRIMARY KEY (`idMonster`),
  UNIQUE INDEX `Name_UNIQUE` (`Name` ASC))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `PathfinderEncounter`.`Attacks`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `PathfinderEncounter`.`Attacks` (
  `idAttacks` INT NOT NULL,
  `AttackName` VARCHAR(45) NULL,
  `-4` VARCHAR(45) NOT NULL,
  `-3` VARCHAR(45) NOT NULL,
  `-2` VARCHAR(45) NOT NULL,
  `-1` VARCHAR(45) NOT NULL,
  `0` VARCHAR(45) NOT NULL,
  `1` VARCHAR(45) NOT NULL,
  `2` VARCHAR(45) NOT NULL,
  `3` VARCHAR(45) NOT NULL,
  `4` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`idAttacks`),
  UNIQUE INDEX `Attack Name_UNIQUE` (`AttackName` ASC),
  CONSTRAINT `Size`
    FOREIGN KEY ()
    REFERENCES `PathfinderEncounter`.`Monster` ()
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `Str`
    FOREIGN KEY ()
    REFERENCES `PathfinderEncounter`.`Monster` ()
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `Dex`
    FOREIGN KEY ()
    REFERENCES `PathfinderEncounter`.`Monster` ()
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `PathfinderEncounter`.`Type`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `PathfinderEncounter`.`Type` (
  `idType` INT NOT NULL,
  `TypeName` VARCHAR(45) NOT NULL,
  `HitDie` INT NOT NULL,
  PRIMARY KEY (`idType`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `PathfinderEncounter`.`Monster_has_Type`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `PathfinderEncounter`.`Monster_has_Type` (
  `Monster_idMonster` INT NOT NULL,
  `Type_idType` INT NOT NULL,
  PRIMARY KEY (`Monster_idMonster`, `Type_idType`),
  INDEX `fk_Monster_has_Type_Type1_idx` (`Type_idType` ASC),
  INDEX `fk_Monster_has_Type_Monster1_idx` (`Monster_idMonster` ASC),
  CONSTRAINT `fk_Monster_has_Type_Monster1`
    FOREIGN KEY (`Monster_idMonster`)
    REFERENCES `PathfinderEncounter`.`Monster` (`idMonster`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_Monster_has_Type_Type1`
    FOREIGN KEY (`Type_idType`)
    REFERENCES `PathfinderEncounter`.`Type` (`idType`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `PathfinderEncounter`.`Monster_has_Attacks`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `PathfinderEncounter`.`Monster_has_Attacks` (
  `Monster_idMonster` INT NOT NULL,
  `Attacks_idAttacks` INT NOT NULL,
  PRIMARY KEY (`Monster_idMonster`, `Attacks_idAttacks`),
  INDEX `fk_Monster_has_Attacks_Attacks1_idx` (`Attacks_idAttacks` ASC),
  INDEX `fk_Monster_has_Attacks_Monster1_idx` (`Monster_idMonster` ASC),
  CONSTRAINT `fk_Monster_has_Attacks_Monster1`
    FOREIGN KEY (`Monster_idMonster`)
    REFERENCES `PathfinderEncounter`.`Monster` (`idMonster`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_Monster_has_Attacks_Attacks1`
    FOREIGN KEY (`Attacks_idAttacks`)
    REFERENCES `PathfinderEncounter`.`Attacks` (`idAttacks`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;

USE `PathfinderEncounter`;

DELIMITER $$
USE `PathfinderEncounter`$$
CREATE DEFINER = DatabaseManager TRIGGER `PathfinderEncounter`.`Monster_BEFORE_INSERT` BEFORE INSERT ON `Monster` 
FOR EACH ROW
Begin
	DECLARE msg VARCHAR(255);
    DECLARE found_it INT;
    IF NEW.size < -4 OR NEW.size > 4 THEN
		set msg = "Error: Size categories must be an integer between -4 and 4";
        SIGNAL SQLSTATE '45000' SET MESSAGE_TEXT = msg;
        END IF;
        
	IF NEW.alignment <> 'LG' OR NEW.alignment <> 'NG' OR NEW.alignment <> 'CG' OR NEW.alignment <> 'LN' OR NEW.alignment <> 'N' OR NEW.alignment <> 'CN' OR NEW.alignment <> 'LE' OR NEW.alignment <> 'NE' OR NEW.alignment <> 'CE' Then
		set msg = "Error: Invalid input for alignment";
        SIGNAL SQLSTATE '45000' SET MESSAGE_TEXT = msg;
        END IF;
        
    IF NEW.CR < 1 OR NEW.CR > 10 THEN
		set msg = "Error: Size categories must be an integer between -4 and 4";
        SIGNAL SQLSTATE '45000' SET MESSAGE_TEXT = msg;
        END IF;   
        
    IF NEW.Armor < 0 OR NEW.Shield < 0 OR New.Deflection < 0 OR NEW.NaturalArmor < 0 OR NEW.Dodge < 0 OR New.MiscAC < 0 Then   
        set msg = "Error: All Armor Class values must be greater than 0";
        SIGNAL SQLSTATE '45000' SET MESSAGE_TEXT = msg;
        END IF;  
        
    IF NEW.HitDie < 0 OR NEW.BaseSpeed < 0 OR New.Space < 0 OR NEW.Reach < 0 Then   
        set msg = "Error: Hit Die, Speed, Space, and Reach must be nonnegative";
        SIGNAL SQLSTATE '45000' SET MESSAGE_TEXT = msg;
        END IF;  
        
      IF NEW.STR < 0 OR NEW.DEX < 0 OR New.CON < 0 OR NEW.INTE < 0 OR New.WIS < 0 OR NEW.CHA < 0 Then   
        set msg = "Error: Stats must be nonnegative";
        SIGNAL SQLSTATE '45000' SET MESSAGE_TEXT = msg;
        END IF;
        
      SELECT COUNT(1) INTO found_it FROM Attacks.AttackName
        WHERE Attack1 = NEW.Attack1;
        IF found_it = 0 THEN
			set msg = "Error: Attack type not in database. Contact Database Manager for Assistance";
			SIGNAL SQLSTATE '45000' SET MESSAGE_TEXT = msg; 
        END IF;
        
	  SELECT COUNT(1) INTO found_it FROM Attacks.AttackName
        WHERE Attack2 = NEW.Attack2;
        IF found_it = 0 THEN
			set msg = "Error: Attack type not in database. Contact Database Manager for Assistance";
			SIGNAL SQLSTATE '45000' SET MESSAGE_TEXT = msg; 
        END IF;
        
      SELECT COUNT(1) INTO found_it FROM Attacks.AttackName
        WHERE Attack3 = NEW.Attack3;
        IF found_it = 0 THEN
			set msg = "Error: Attack type not in database. Contact Database Manager for Assistance";
			SIGNAL SQLSTATE '45000' SET MESSAGE_TEXT = msg; 
        END IF;
      
      SELECT COUNT(1) INTO found_it FROM Attacks.AttackName
        WHERE Attack4 = NEW.Attack4;
        IF found_it = 0 THEN
			set msg = "Error: Attack type not in database. Contact Database Manager for Assistance";
			SIGNAL SQLSTATE '45000' SET MESSAGE_TEXT = msg; 
        END IF;
        
  	  SELECT COUNT(1) INTO found_it FROM Attacks.AttackName
        WHERE Attack5 = NEW.Attack5;
        IF found_it = 0 THEN
			set msg = "Error: Attack type not in database. Contact Database Manager for Assistance";
			SIGNAL SQLSTATE '45000' SET MESSAGE_TEXT = msg; 
        END IF;      
        END;$$

USE `PathfinderEncounter`$$
CREATE DEFINER = DatabaseManager TRIGGER `PathfinderEncounter`.`Monster_BEFORE_UPDATE` BEFORE UPDATE ON `Monster` FOR EACH ROW
    Begin
	DECLARE msg VARCHAR(255);
    DECLARE found_it INT;
    IF NEW.size < -4 OR NEW.size > 4 THEN
		set msg = "Error: Size categories must be an integer between -4 and 4";
        SIGNAL SQLSTATE '45000' SET MESSAGE_TEXT = msg;
        END IF;
        
	IF NEW.alignment <> 'LG' OR NEW.alignment <> 'NG' OR NEW.alignment <> 'CG' OR NEW.alignment <> 'LN' OR NEW.alignment <> 'N' OR NEW.alignment <> 'CN' OR NEW.alignment <> 'LE' OR NEW.alignment <> 'NE' OR NEW.alignment <> 'CE' Then
		set msg = "Error: Invalid input for alignment";
        SIGNAL SQLSTATE '45000' SET MESSAGE_TEXT = msg;
        END IF;
        
    IF NEW.CR < 1 OR NEW.CR > 10 THEN
		set msg = "Error: Size categories must be an integer between -4 and 4";
        SIGNAL SQLSTATE '45000' SET MESSAGE_TEXT = msg;
        END IF;   
        
    IF NEW.Armor < 0 OR NEW.Shield < 0 OR New.Deflection < 0 OR NEW.NaturalArmor < 0 OR NEW.Dodge < 0 OR New.MiscAC < 0 Then   
        set msg = "Error: All Armor Class values must be greater than 0";
        SIGNAL SQLSTATE '45000' SET MESSAGE_TEXT = msg;
        END IF;  
        
    IF NEW.HitDie < 0 OR NEW.BaseSpeed < 0 OR New.Space < 0 OR NEW.Reach < 0 Then   
        set msg = "Error: Hit Die, Speed, Space, and Reach must be nonnegative";
        SIGNAL SQLSTATE '45000' SET MESSAGE_TEXT = msg;
        END IF;  
        
      IF NEW.STR < 0 OR NEW.DEX < 0 OR New.CON < 0 OR NEW.INTE < 0 OR New.WIS < 0 OR NEW.CHA < 0 Then   
        set msg = "Error: Stats must be nonnegative";
        SIGNAL SQLSTATE '45000' SET MESSAGE_TEXT = msg;
        END IF;
        
      SELECT COUNT(1) INTO found_it FROM Attacks.AttackName
        WHERE Attack1 = NEW.Attack1;
        IF found_it = 0 THEN
			set msg = "Error: Attack type not in database. Contact Database Manager for Assistance";
			SIGNAL SQLSTATE '45000' SET MESSAGE_TEXT = msg; 
        END IF;
        
	  SELECT COUNT(1) INTO found_it FROM Attacks.AttackName
        WHERE Attack2 = NEW.Attack2;
        IF found_it = 0 THEN
			set msg = "Error: Attack type not in database. Contact Database Manager for Assistance";
			SIGNAL SQLSTATE '45000' SET MESSAGE_TEXT = msg; 
        END IF;
        
      SELECT COUNT(1) INTO found_it FROM Attacks.AttackName
        WHERE Attack3 = NEW.Attack3;
        IF found_it = 0 THEN
			set msg = "Error: Attack type not in database. Contact Database Manager for Assistance";
			SIGNAL SQLSTATE '45000' SET MESSAGE_TEXT = msg; 
        END IF;
      
      SELECT COUNT(1) INTO found_it FROM Attacks.AttackName
        WHERE Attack4 = NEW.Attack4;
        IF found_it = 0 THEN
			set msg = "Error: Attack type not in database. Contact Database Manager for Assistance";
			SIGNAL SQLSTATE '45000' SET MESSAGE_TEXT = msg; 
        END IF;
        
  	  SELECT COUNT(1) INTO found_it FROM Attacks.AttackName
        WHERE Attack5 = NEW.Attack5;
        IF found_it = 0 THEN
			set msg = "Error: Attack type not in database. Contact Database Manager for Assistance";
			SIGNAL SQLSTATE '45000' SET MESSAGE_TEXT = msg; 
        END IF;      
        END;$$

USE `PathfinderEncounter`$$
CREATE DEFINER = CURRENT_USER TRIGGER `PathfinderEncounter`.`Attacks_BEFORE_DELETE` BEFORE DELETE ON `Attacks` FOR EACH ROW
   begin
 	DECLARE msg VARCHAR(255);
    DECLARE found_it INT;
    
       SELECT COUNT(1) INTO found_it FROM Monster.Attack1
        WHERE AttackName = NEW.AttackName;
        IF found_it = 0 THEN
			set msg = "Error: Monster in database has attack and must be deleted before attack can be removed";
			SIGNAL SQLSTATE '45000' SET MESSAGE_TEXT = msg; 
        END IF;
        
       SELECT COUNT(1) INTO found_it FROM Monster.Attack2
        WHERE AttackName = NEW.AttackName;
        IF found_it = 0 THEN
			set msg = "Error: Monster in database has attack and must be deleted before attack can be removed";
			SIGNAL SQLSTATE '45000' SET MESSAGE_TEXT = msg; 
        END IF;
        
       SELECT COUNT(1) INTO found_it FROM Monster.Attack3
        WHERE AttackName = NEW.AttackName;
        IF found_it = 0 THEN
			set msg = "Error: Monster in database has attack and must be deleted before attack can be removed";
			SIGNAL SQLSTATE '45000' SET MESSAGE_TEXT = msg; 
        END IF; 
        
       SELECT COUNT(1) INTO found_it FROM Monster.Attack4
        WHERE AttackName = NEW.AttackName;
        IF found_it = 0 THEN
			set msg = "Error: Monster in database has attack and must be deleted before attack can be removed";
			SIGNAL SQLSTATE '45000' SET MESSAGE_TEXT = msg; 
        END IF;      
        
       SELECT COUNT(1) INTO found_it FROM Monster.Attack5
        WHERE AttackName = NEW.AttackName;
        IF found_it = 0 THEN
			set msg = "Error: Monster in database has attack and must be deleted before attack can be removed";
			SIGNAL SQLSTATE '45000' SET MESSAGE_TEXT = msg; 
        END IF;     
        End;  $$

USE `PathfinderEncounter`$$
CREATE DEFINER = CURRENT_USER TRIGGER `PathfinderEncounter`.`Type_BEFORE_INSERT` BEFORE INSERT ON `Type` FOR EACH ROW
Begin
	DECLARE msg VARCHAR(255);
    
    IF NEW.HitDie < 1 THEN
		set msg = "Error: Size of Hit Die must be greater than 0";
        SIGNAL SQLSTATE '45000' SET MESSAGE_TEXT = msg;
        END IF;
        END;
            $$

USE `PathfinderEncounter`$$
CREATE DEFINER = CURRENT_USER TRIGGER `PathfinderEncounter`.`Type_BEFORE_DELETE` BEFORE DELETE ON `Type` FOR EACH ROW
 begin
 	DECLARE msg VARCHAR(255);
    DECLARE found_it INT;
    
       SELECT COUNT(1) INTO found_it FROM Monster.TypeName
        WHERE TypeName = NEW.TypeName;
        IF found_it = 0 THEN
			set msg = "Error: Monster in database has type and must be deleted before type can be removed";
			SIGNAL SQLSTATE '45000' SET MESSAGE_TEXT = msg; 
        END IF;
        End;$$


DELIMITER ;
CREATE USER 'DatabaseManager' IDENTIFIED BY 'database';

GRANT ALL ON `PathfinderEncounter`.* TO 'DatabaseManager';
GRANT SELECT ON TABLE `PathfinderEncounter`.* TO 'DatabaseManager';
GRANT SELECT, INSERT, TRIGGER ON TABLE `PathfinderEncounter`.* TO 'DatabaseManager';
GRANT SELECT, INSERT, TRIGGER, UPDATE, DELETE ON TABLE `PathfinderEncounter`.* TO 'DatabaseManager';
GRANT EXECUTE ON ROUTINE `PathfinderEncounter`.* TO 'DatabaseManager';
CREATE USER 'User';

GRANT SELECT ON TABLE `PathfinderEncounter`.* TO 'User';
CREATE USER 'Engineer';

GRANT SELECT, INSERT, TRIGGER ON TABLE `PathfinderEncounter`.* TO 'Engineer';
GRANT SELECT, INSERT, TRIGGER, UPDATE, DELETE ON TABLE `PathfinderEncounter`.* TO 'Engineer';
GRANT EXECUTE ON ROUTINE `PathfinderEncounter`.* TO 'Engineer';

SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;