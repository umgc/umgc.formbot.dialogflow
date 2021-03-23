CREATE TABLE `form` (
  `ID` bigint(20) NOT NULL,
  `name` varchar(250) NOT NULL,
  `description` varchar(3000) DEFAULT NULL,
  PRIMARY KEY (`ID`),
  UNIQUE KEY `name_UNIQUE` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

;


CREATE TABLE `form_entry` (
  `ID` bigint(20) NOT NULL,
  `formId` bigint(20) DEFAULT NULL,
  `userId` bigint(20) DEFAULT NULL,
  `created` varchar(45) DEFAULT 'CURRENT_TIMESTAMP',
  `modified` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  KEY `fk_formId_idx` (`ID`),
  KEY `fk_form_id_idx` (`formId`),
  KEY `fk_user_idx` (`userId`),
  CONSTRAINT `fk_form_id` FOREIGN KEY (`formId`) REFERENCES `form` (`ID`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `fk_user` FOREIGN KEY (`userId`) REFERENCES `user` (`Id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8


;

CREATE TABLE `form_entry_field` (
  `ID` bigint(20) NOT NULL,
  `entryFormFieldId` bigint(20) DEFAULT NULL,
  `entryFormId` bigint(20) DEFAULT NULL,
  `entry` varchar(6000) DEFAULT NULL,
  `created` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `modified` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`ID`),
  KEY `fk_entry_form_id_idx` (`entryFormId`),
  CONSTRAINT `fk_entry_form_id` FOREIGN KEY (`entryFormId`) REFERENCES `form_field` (`ID`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8

;

CREATE TABLE `form_entry_field_map` (
  `ID` bigint(20) NOT NULL,
  `session` varchar(250) NOT NULL,
  `formEntryId` bigint(20) DEFAULT NULL,
  `formEntryFieldId` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`ID`),
  KEY `fk_form_entry_field_id_idx` (`formEntryFieldId`),
  KEY `fk_form_entry_id_idx` (`formEntryId`),
  CONSTRAINT `fk_form_entry_field_id` FOREIGN KEY (`formEntryFieldId`) REFERENCES `form_entry_field` (`ID`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `fk_form_entry_id` FOREIGN KEY (`formEntryId`) REFERENCES `form_entry` (`ID`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8

;

CREATE TABLE `form_field` (
  `ID` bigint(20) NOT NULL,
  `name` varchar(150) NOT NULL,
  `description` varchar(6000) DEFAULT NULL,
  PRIMARY KEY (`ID`),
  UNIQUE KEY `name_UNIQUE` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

;

CREATE TABLE `form_field_map` (
  `ID` bigint(20) NOT NULL,
  `formId` bigint(20) NOT NULL,
  `formFieldId` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`ID`),
  KEY `fk_formId_idx` (`formId`),
  KEY `fk_formfieldId_idx` (`formFieldId`),
  CONSTRAINT `fk_formId` FOREIGN KEY (`formId`) REFERENCES `form` (`ID`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `fk_formfieldId` FOREIGN KEY (`formFieldId`) REFERENCES `form_field` (`ID`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8

;

CREATE TABLE `user` (
  `Id` bigint(20) NOT NULL,
  `nameFirst` varchar(150) DEFAULT NULL,
  `nameLast` varchar(150) DEFAULT NULL,
  `roll` varchar(45) DEFAULT NULL,
  `email` varchar(250) DEFAULT NULL,
  `password` varchar(6000) DEFAULT NULL,
  `created` datetime DEFAULT CURRENT_TIMESTAMP,
  `modified` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8