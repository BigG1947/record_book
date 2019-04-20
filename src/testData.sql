DELETE from marks;
ALTER TABLE marks AUTO_INCREMENT=0;
DELETE FROM loads;
ALTER TABLE loads AUTO_INCREMENT=0;
DELETE FROM discipline;
ALTER TABLE discipline AUTO_INCREMENT=0;
DELETE FROM sensitive_data;
ALTER TABLE sensitive_data AUTO_INCREMENT=0;
DELETE FROM student;
ALTER TABLE student AUTO_INCREMENT=0;
DELETE FROM employee;
ALTER TABLE employee AUTO_INCREMENT=0;
DELETE FROM accession;
ALTER TABLE accession AUTO_INCREMENT=0;
DELETE FROM people;
ALTER TABLE people AUTO_INCREMENT=0;
DELETE FROM speciality;
ALTER TABLE speciality AUTO_INCREMENT=0;
DELETE FROM direction;
ALTER TABLE direction AUTO_INCREMENT=0;
DELETE FROM cathedra;
ALTER TABLE cathedra AUTO_INCREMENT=0;
DELETE FROM faculty;
ALTER TABLE faculty AUTO_INCREMENT=0;
DELETE FROM institute;
ALTER TABLE institute AUTO_INCREMENT=0;
-- DELETE FROM ranks;
-- ALTER TABLE ranks AUTO_INCREMENT=0;
DELETE FROM groups;
ALTER TABLE groups AUTO_INCREMENT=0;
-- DELETE FROM status;
-- ALTER TABLE status AUTO_INCREMENT=0;


INSERT INTO institute(name) VALUES ("institute1");
INSERT INTO institute(name) VALUES ("institute2");

INSERT INTO faculty(name, id_institute) VALUES ("faculty1", 1);
INSERT INTO faculty(name, id_institute) VALUES ("faculty2", 1);
INSERT INTO faculty(name, id_institute) VALUES ("faculty3", 2);
INSERT INTO faculty(name, id_institute) VALUES ("faculty4", 2);

INSERT INTO cathedra(name, id_faculty) VALUES ("cathedra1", 1);
INSERT INTO cathedra(name, id_faculty) VALUES ("cathedra2", 1);
INSERT INTO cathedra(name, id_faculty) VALUES ("cathedra3", 2);
INSERT INTO cathedra(name, id_faculty) VALUES ("cathedra4", 2);
INSERT INTO cathedra(name, id_faculty) VALUES ("cathedra5", 3);
INSERT INTO cathedra(name, id_faculty) VALUES ("cathedra6", 3);
INSERT INTO cathedra(name, id_faculty) VALUES ("cathedra7", 4);
INSERT INTO cathedra(name, id_faculty) VALUES ("cathedra8", 4);

INSERT INTO direction(name, id_cathedra) VALUES ("direction1", 1);
INSERT INTO direction(name, id_cathedra) VALUES ("direction2", 1);
INSERT INTO direction(name, id_cathedra) VALUES ("direction3", 2);
INSERT INTO direction(name, id_cathedra) VALUES ("direction4", 2);
INSERT INTO direction(name, id_cathedra) VALUES ("direction5", 3);
INSERT INTO direction(name, id_cathedra) VALUES ("direction6", 3);
INSERT INTO direction(name, id_cathedra) VALUES ("direction7", 4);
INSERT INTO direction(name, id_cathedra) VALUES ("direction8", 4);
INSERT INTO direction(name, id_cathedra) VALUES ("direction9", 5);
INSERT INTO direction(name, id_cathedra) VALUES ("direction10", 5);
INSERT INTO direction(name, id_cathedra) VALUES ("direction11", 6);
INSERT INTO direction(name, id_cathedra) VALUES ("direction12", 6);
INSERT INTO direction(name, id_cathedra) VALUES ("direction13", 7);
INSERT INTO direction(name, id_cathedra) VALUES ("direction14", 7);
INSERT INTO direction(name, id_cathedra) VALUES ("direction15", 8);
INSERT INTO direction(name, id_cathedra) VALUES ("direction16", 8);

INSERT INTO speciality(name, id_direction) VALUES ("speciality1", 1);
INSERT INTO speciality(name, id_direction) VALUES ("speciality2", 1);
INSERT INTO speciality(name, id_direction) VALUES ("speciality3", 2);
INSERT INTO speciality(name, id_direction) VALUES ("speciality4", 2);
INSERT INTO speciality(name, id_direction) VALUES ("speciality5", 3);
INSERT INTO speciality(name, id_direction) VALUES ("speciality6", 3);
INSERT INTO speciality(name, id_direction) VALUES ("speciality7", 4);
INSERT INTO speciality(name, id_direction) VALUES ("speciality8", 4);
INSERT INTO speciality(name, id_direction) VALUES ("speciality9", 5);
INSERT INTO speciality(name, id_direction) VALUES ("speciality10", 5);
INSERT INTO speciality(name, id_direction) VALUES ("speciality11", 6);
INSERT INTO speciality(name, id_direction) VALUES ("speciality12", 6);
INSERT INTO speciality(name, id_direction) VALUES ("speciality13", 7);
INSERT INTO speciality(name, id_direction) VALUES ("speciality14", 7);
INSERT INTO speciality(name, id_direction) VALUES ("speciality15", 8);
INSERT INTO speciality(name, id_direction) VALUES ("speciality16", 8);
INSERT INTO speciality(name, id_direction) VALUES ("speciality17", 9);
INSERT INTO speciality(name, id_direction) VALUES ("speciality18", 9);
INSERT INTO speciality(name, id_direction) VALUES ("speciality19", 10);
INSERT INTO speciality(name, id_direction) VALUES ("speciality20", 10);
INSERT INTO speciality(name, id_direction) VALUES ("speciality21", 11);
INSERT INTO speciality(name, id_direction) VALUES ("speciality22", 11);
INSERT INTO speciality(name, id_direction) VALUES ("speciality23", 12);
INSERT INTO speciality(name, id_direction) VALUES ("speciality24", 12);
INSERT INTO speciality(name, id_direction) VALUES ("speciality25", 13);
INSERT INTO speciality(name, id_direction) VALUES ("speciality26", 13);
INSERT INTO speciality(name, id_direction) VALUES ("speciality27", 14);
INSERT INTO speciality(name, id_direction) VALUES ("speciality28", 14);
INSERT INTO speciality(name, id_direction) VALUES ("speciality29", 15);
INSERT INTO speciality(name, id_direction) VALUES ("speciality30", 15);
INSERT INTO speciality(name, id_direction) VALUES ("speciality31", 16);
INSERT INTO speciality(name, id_direction) VALUES ("speciality32", 16);

-- INSERT INTO ranks(name) VALUES ("rank1");
-- INSERT INTO ranks(name) VALUES ("rank2");
-- INSERT INTO ranks(name) VALUES ("rank3");
-- INSERT INTO ranks(name) VALUES ("rank4");
-- INSERT INTO ranks(name) VALUES ("rank5");
-- INSERT INTO ranks(name) VALUES ("rank6");
--
-- INSERT INTO status(name) VALUES ("status1");
-- INSERT INTO status(name) VALUES ("status2");
-- INSERT INTO status(name) VALUES ("status3");
-- INSERT INTO status(name) VALUES ("status4");
-- INSERT INTO status(name) VALUES ("status5");

INSERT INTO people(fio, birthday, gender, img, comment, password, phone_number, email, id_status, have_access)
VALUES ("Степул Артем Мартиросовчи", "1997-06-12", 1, "/static/img/default.png", "Поет, учавствует в студ. клубе", "$2a$10$oqygBXUWPKkdL.wdac68n.GtjRGxk8LnaqxpTcNvZGWxRTtRHE0TK", "+380999999999", "artem.stepul@gmail.com", 2, 1);
INSERT INTO people(fio, birthday, gender, img, comment, password, phone_number, email, id_status, have_access)
VALUES ("Луценко Артем Геннадиевич", "1997-03-16", 1, "/static/img/default.png", "Заместитель главы студенчиского совета факультета", "$2a$10$oqygBXUWPKkdL.wdac68n.GtjRGxk8LnaqxpTcNvZGWxRTtRHE0TK", "+380333333333", "artem.lucenko@gmail.com", 2, 1);
INSERT INTO people(fio, birthday, gender, img, comment, password, phone_number, email, id_status, have_access)
VALUES ("Иванов Иван Иванович", "1979-09-17", 1, "/static/img/default.png", "", "$2a$10$oqygBXUWPKkdL.wdac68n.GtjRGxk8LnaqxpTcNvZGWxRTtRHE0TK", "+380444444444", "admin@admin.com", 1, 1);

INSERT INTO people(fio, birthday, gender, img, comment, password, phone_number, email, id_status, have_access)
VALUES ("st3", "1997-06-12", 1, "/static/img/default.png", "Поет, учавствует в студ. клубе", "$2a$10$oqygBXUWPKkdL.wdac68n.GtjRGxk8LnaqxpTcNvZGWxRTtRHE0TK", "+380999999999", "st3@gmail.com", 2, 1);
INSERT INTO people(fio, birthday, gender, img, comment, password, phone_number, email, id_status, have_access)
VALUES ("st4", "1997-03-16", 1, "/static/img/default.png", "Заместитель главы студенчиского совета факультета", "$2a$10$oqygBXUWPKkdL.wdac68n.GtjRGxk8LnaqxpTcNvZGWxRTtRHE0TK", "+380333333333", "st4@gmail.com", 2, 1);
INSERT INTO people(fio, birthday, gender, img, comment, password, phone_number, email, id_status, have_access)
VALUES ("empl2", "1979-09-17", 1, "/static/img/default.png", "", "$2a$10$oqygBXUWPKkdL.wdac68n.GtjRGxk8LnaqxpTcNvZGWxRTtRHE0TK", "+380444444444", "empl2@admin.com", 1, 1);
INSERT INTO people(fio, birthday, gender, img, comment, password, phone_number, email, id_status, have_access)
VALUES ("st5", "1997-06-12", 1, "/static/img/default.png", "Поет, учавствует в студ. клубе", "$2a$10$oqygBXUWPKkdL.wdac68n.GtjRGxk8LnaqxpTcNvZGWxRTtRHE0TK", "+380999999999", "st5@gmail.com", 2, 1);
INSERT INTO people(fio, birthday, gender, img, comment, password, phone_number, email, id_status, have_access)
VALUES ("st6", "1997-03-16", 1, "/static/img/default.png", "Заместитель главы студенчиского совета факультета", "$2a$10$oqygBXUWPKkdL.wdac68n.GtjRGxk8LnaqxpTcNvZGWxRTtRHE0TK", "+380333333333", "st6@gmail.com", 2, 1);
INSERT INTO people(fio, birthday, gender, img, comment, password, phone_number, email, id_status, have_access)
VALUES ("empl3", "1979-09-17", 1, "/static/img/default.png", "", "$2a$10$oqygBXUWPKkdL.wdac68n.GtjRGxk8LnaqxpTcNvZGWxRTtRHE0TK", "+380444444444", "empl3@admin.com", 1, 1);
INSERT INTO people(fio, birthday, gender, img, comment, password, phone_number, email, id_status, have_access)
VALUES ("st7", "1997-06-12", 1, "/static/img/default.png", "Поет, учавствует в студ. клубе", "$2a$10$oqygBXUWPKkdL.wdac68n.GtjRGxk8LnaqxpTcNvZGWxRTtRHE0TK", "+380999999999", "st7@gmail.com", 2, 1);
INSERT INTO people(fio, birthday, gender, img, comment, password, phone_number, email, id_status, have_access)
VALUES ("st8", "1997-03-16", 1, "/static/img/default.png", "Заместитель главы студенчиского совета факультета", "$2a$10$oqygBXUWPKkdL.wdac68n.GtjRGxk8LnaqxpTcNvZGWxRTtRHE0TK", "+380333333333", "st8@gmail.com", 2, 1);
INSERT INTO people(fio, birthday, gender, img, comment, password, phone_number, email, id_status, have_access)
VALUES ("empl4", "1979-09-17", 1, "/static/img/default.png", "", "$2a$10$oqygBXUWPKkdL.wdac68n.GtjRGxk8LnaqxpTcNvZGWxRTtRHE0TK", "+380444444444", "empl4@admin.com", 1, 1);
INSERT INTO people(fio, birthday, gender, img, comment, password, phone_number, email, id_status, have_access)
VALUES ("st9", "1997-06-12", 1, "/static/img/default.png", "Поет, учавствует в студ. клубе", "$2a$10$oqygBXUWPKkdL.wdac68n.GtjRGxk8LnaqxpTcNvZGWxRTtRHE0TK", "+380999999999", "st9@gmail.com", 2, 1);
INSERT INTO people(fio, birthday, gender, img, comment, password, phone_number, email, id_status, have_access)
VALUES ("st10", "1997-03-16", 1, "/static/img/default.png", "Заместитель главы студенчиского совета факультета", "$2a$10$oqygBXUWPKkdL.wdac68n.GtjRGxk8LnaqxpTcNvZGWxRTtRHE0TK", "+380333333333", "st10@gmail.com", 2, 1);

INSERT INTO employee(id_people, date_invite, id_cathedra, id_rank)
VALUES (3, "2007-07-22", 1, 1);
INSERT INTO employee(id_people, date_invite, id_cathedra, id_rank)
VALUES (6, "2007-07-22", 1, 1);
INSERT INTO employee(id_people, date_invite, id_cathedra, id_rank)
VALUES (9, "2007-07-22", 1, 1);
INSERT INTO employee(id_people, date_invite, id_cathedra, id_rank)
VALUES (12, "2007-07-22", 1, 1);


INSERT INTO groups(id_employee, name, id_direction) VALUES (3, "group1", 1);
INSERT INTO groups(id_employee, name, id_direction) VALUES (NULL, "group2", 1);
INSERT INTO groups(id_employee, name, id_direction) VALUES (NULL, "group3", 2);
INSERT INTO groups(id_employee, name, id_direction) VALUES (NULL, "group4", 2);
INSERT INTO groups(id_employee, name, id_direction) VALUES (NULL, "group5", 3);
INSERT INTO groups(id_employee, name, id_direction) VALUES (NULL, "group6", 3);
INSERT INTO groups(id_employee, name, id_direction) VALUES (NULL, "group7", 4);
INSERT INTO groups(id_employee, name, id_direction) VALUES (NULL, "group8", 4);
INSERT INTO groups(id_employee, name, id_direction) VALUES (NULL, "group9", 5);
INSERT INTO groups(id_employee, name, id_direction) VALUES (NULL, "group10", 5);
INSERT INTO groups(id_employee, name, id_direction) VALUES (NULL, "group11", 6);
INSERT INTO groups(id_employee, name, id_direction) VALUES (NULL, "group12", 6);
INSERT INTO groups(id_employee, name, id_direction) VALUES (NULL, "group13", 7);
INSERT INTO groups(id_employee, name, id_direction) VALUES (NULL, "group14", 7);
INSERT INTO groups(id_employee, name, id_direction) VALUES (NULL, "group15", 8);
INSERT INTO groups(id_employee, name, id_direction) VALUES (NULL, "group16", 8);
INSERT INTO groups(id_employee, name, id_direction) VALUES (NULL, "group17", 9);
INSERT INTO groups(id_employee, name, id_direction) VALUES (NULL, "group18", 9);
INSERT INTO groups(id_employee, name, id_direction) VALUES (3, "group19", 10);
INSERT INTO groups(id_employee, name, id_direction) VALUES (NULL, "group20", 10);
INSERT INTO groups(id_employee, name, id_direction) VALUES (NULL, "group21", 11);
INSERT INTO groups(id_employee, name, id_direction) VALUES (NULL, "group22", 11);
INSERT INTO groups(id_employee, name, id_direction) VALUES (NULL, "group23", 12);
INSERT INTO groups(id_employee, name, id_direction) VALUES (NULL, "group24", 12);
INSERT INTO groups(id_employee, name, id_direction) VALUES (NULL, "group25", 13);
INSERT INTO groups(id_employee, name, id_direction) VALUES (NULL, "group26", 13);
INSERT INTO groups(id_employee, name, id_direction) VALUES (NULL, "group27", 14);
INSERT INTO groups(id_employee, name, id_direction) VALUES (NULL, "group28", 14);
INSERT INTO groups(id_employee, name, id_direction) VALUES (NULL, "group29", 15);
INSERT INTO groups(id_employee, name, id_direction) VALUES (NULL, "group30", 15);
INSERT INTO groups(id_employee, name, id_direction) VALUES (NULL, "group31", 16);
INSERT INTO groups(id_employee, name, id_direction) VALUES (NULL, "group32", 16);

INSERT INTO student(id_people, date_admission, is_full_time, is_cut, semester, id_group)
VALUES (1, "2017-09-01", 1, 1, 4, 1);
INSERT INTO student(id_people, date_admission, is_full_time, is_cut, semester, id_group)
VALUES (2, "2017-09-01", 1, 1, 4, 1);
INSERT INTO student(id_people, date_admission, is_full_time, is_cut, semester, id_group)
VALUES (4, "2017-09-01", 1, 1, 4, 1);
INSERT INTO student(id_people, date_admission, is_full_time, is_cut, semester, id_group)
VALUES (5, "2017-09-01", 1, 1, 4, 1);
INSERT INTO student(id_people, date_admission, is_full_time, is_cut, semester, id_group)
VALUES (7, "2017-09-01", 1, 1, 4, 1);
INSERT INTO student(id_people, date_admission, is_full_time, is_cut, semester, id_group)
VALUES (8, "2017-09-01", 1, 1, 4, 1);
INSERT INTO student(id_people, date_admission, is_full_time, is_cut, semester, id_group)
VALUES (10, "2017-09-01", 1, 1, 4, 2);
INSERT INTO student(id_people, date_admission, is_full_time, is_cut, semester, id_group)
VALUES (11, "2017-09-01", 1, 1, 4, 2);
INSERT INTO student(id_people, date_admission, is_full_time, is_cut, semester, id_group)
VALUES (13, "2017-09-01", 1, 1, 4, 2);
INSERT INTO student(id_people, date_admission, is_full_time, is_cut, semester, id_group)
VALUES (14, "2017-09-01", 1, 1, 4, 2);



INSERT INTO accession(id_people, edit_access, set_absence, get_absence, set_mark, set_event, get_sensitive, set_sensitive, get_ylist, manage_academ)
VALUES (1, 0, 0, 0, 0, 0, 0, 0, 0, 0);
INSERT INTO accession(id_people, edit_access, set_absence, get_absence, set_mark, set_event, get_sensitive, set_sensitive, get_ylist, manage_academ)
VALUES (2, 1, 0, 1, 0, 0, 1, 1, 0, 1);
INSERT INTO accession(id_people, edit_access, set_absence, get_absence, set_mark, set_event, get_sensitive, set_sensitive, get_ylist, manage_academ)
VALUES (3, 1, 1, 1, 1, 1, 1, 1, 1, 1);
INSERT INTO accession(id_people, edit_access, set_absence, get_absence, set_mark, set_event, get_sensitive, set_sensitive, get_ylist, manage_academ)
VALUES (4, 1, 0, 1, 0, 1, 0, 1, 0, 1);
INSERT INTO accession(id_people, edit_access, set_absence, get_absence, set_mark, set_event, get_sensitive, set_sensitive, get_ylist, manage_academ)
VALUES (5, 0, 1, 0, 1, 0, 1, 0, 1, 0);
INSERT INTO accession(id_people, edit_access, set_absence, get_absence, set_mark, set_event, get_sensitive, set_sensitive, get_ylist, manage_academ)
VALUES (6, 1, 1, 1, 1, 1, 1, 1, 1, 1);
INSERT INTO accession(id_people, edit_access, set_absence, get_absence, set_mark, set_event, get_sensitive, set_sensitive, get_ylist, manage_academ)
VALUES (7, 0, 1, 0, 1, 0, 1, 0, 1, 0);
INSERT INTO accession(id_people, edit_access, set_absence, get_absence, set_mark, set_event, get_sensitive, set_sensitive, get_ylist, manage_academ)
VALUES (8, 0, 0, 0, 0, 0, 1, 1, 1, 1);
INSERT INTO accession(id_people, edit_access, set_absence, get_absence, set_mark, set_event, get_sensitive, set_sensitive, get_ylist, manage_academ)
VALUES (9, 1, 1, 1, 1, 1, 1, 1, 1, 1);
INSERT INTO accession(id_people, edit_access, set_absence, get_absence, set_mark, set_event, get_sensitive, set_sensitive, get_ylist, manage_academ)
VALUES (10, 1, 1, 1, 1, 0, 0, 0, 0, 0);
INSERT INTO accession(id_people, edit_access, set_absence, get_absence, set_mark, set_event, get_sensitive, set_sensitive, get_ylist, manage_academ)
VALUES (11, 0, 0, 1, 1, 1, 1, 0, 0, 0);
INSERT INTO accession(id_people, edit_access, set_absence, get_absence, set_mark, set_event, get_sensitive, set_sensitive, get_ylist, manage_academ)
VALUES (12, 1, 1, 1, 1, 1, 1, 1, 1, 1);
INSERT INTO accession(id_people, edit_access, set_absence, get_absence, set_mark, set_event, get_sensitive, set_sensitive, get_ylist, manage_academ)
VALUES (13, 0, 0, 1, 1, 1, 0, 0, 1, 1);
INSERT INTO accession(id_people, edit_access, set_absence, get_absence, set_mark, set_event, get_sensitive, set_sensitive, get_ylist, manage_academ)
VALUES (14, 0, 1, 1, 0, 1, 0, 1, 0, 1);

INSERT INTO sensitive_data(id_people, passport_code, rntrs, reg_address, military_id)
VALUES (1, "122fasf2", "214124124", "ул. Пушкина, дом Колатушкина, квартира 25", "ком214124");
INSERT INTO sensitive_data(id_people, passport_code, rntrs, reg_address, military_id)
VALUES (2, "214124124", "122fasf2", "ул. Пушкина, дом Колатушкина, квартира 12", "ыфафаы251512");
INSERT INTO sensitive_data(id_people, passport_code, rntrs, reg_address, military_id)
VALUES (3, "цйкйкf2", "2141фафыафыа", "ул. Пушкина, дом Колатушкина, квартира 36", "пддждхзз215251аыф");
INSERT INTO sensitive_data(id_people, passport_code, rntrs, reg_address, military_id)
VALUES (4, "122fasf2", "214124124", "ул. Пушкина, дом Колатушкина, квартира 25", "ком214124");
INSERT INTO sensitive_data(id_people, passport_code, rntrs, reg_address, military_id)
VALUES (5, "214124124", "122fasf2", "ул. Пушкина, дом Колатушкина, квартира 12", "ыфафаы251512");
INSERT INTO sensitive_data(id_people, passport_code, rntrs, reg_address, military_id)
VALUES (6, "цйкйкf2", "2141фафыафыа", "ул. Пушкина, дом Колатушкина, квартира 36", "пддждхзз215251аыф");
INSERT INTO sensitive_data(id_people, passport_code, rntrs, reg_address, military_id)
VALUES (7, "122fasf2", "214124124", "ул. Пушкина, дом Колатушкина, квартира 25", "ком214124");
INSERT INTO sensitive_data(id_people, passport_code, rntrs, reg_address, military_id)
VALUES (8, "214124124", "122fasf2", "ул. Пушкина, дом Колатушкина, квартира 12", "ыфафаы251512");
INSERT INTO sensitive_data(id_people, passport_code, rntrs, reg_address, military_id)
VALUES (9, "цйкйкf2", "2141фафыафыа", "ул. Пушкина, дом Колатушкина, квартира 36", "пддждхзз215251аыф");
INSERT INTO sensitive_data(id_people, passport_code, rntrs, reg_address, military_id)
VALUES (10, "122fasf2", "214124124", "ул. Пушкина, дом Колатушкина, квартира 25", "ком214124");
INSERT INTO sensitive_data(id_people, passport_code, rntrs, reg_address, military_id)
VALUES (11, "214124124", "122fasf2", "ул. Пушкина, дом Колатушкина, квартира 12", "ыфафаы251512");
INSERT INTO sensitive_data(id_people, passport_code, rntrs, reg_address, military_id)
VALUES (12, "цйкйкf2", "2141фафыафыа", "ул. Пушкина, дом Колатушкина, квартира 36", "пддждхзз215251аыф");
INSERT INTO sensitive_data(id_people, passport_code, rntrs, reg_address, military_id)
VALUES (13, "122fasf2", "214124124", "ул. Пушкина, дом Колатушкина, квартира 25", "ком214124");
INSERT INTO sensitive_data(id_people, passport_code, rntrs, reg_address, military_id)
VALUES (14, "214124124", "122fasf2", "ул. Пушкина, дом Колатушкина, квартира 12", "ыфафаы251512");

INSERT INTO discipline(name) VALUES ("discipline1");
INSERT INTO discipline(name) VALUES ("discipline2");
INSERT INTO discipline(name) VALUES ("discipline3");

INSERT INTO loads(id_discipline, id_employee, id_group, semester, id_assistant) VALUES (1, 3, 1, 4, 6);
INSERT INTO loads(id_discipline, id_employee, id_group, semester, id_assistant) VALUES (2, 3, 1, 4, 6);
INSERT INTO loads(id_discipline, id_employee, id_group, semester, id_assistant) VALUES (1, 6, 2, 4, 6);
INSERT INTO loads(id_discipline, id_employee, id_group, semester, id_assistant) VALUES (2, 6, 2, 4, 6);

INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester)
VALUES (1, 1, 3, 59, "неуд", 1, 4);
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester)
VALUES (2, 1, 3, 59, "неуд", 1, 4);
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester)
VALUES (4, 1, 3, 59, "неуд", 1, 4);
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester)
VALUES (5, 1, 3, 59, "неуд", 1, 4);
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester)
VALUES (7, 1, 3, 59, "неуд", 1, 4);
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester)
VALUES (8, 1, 3, 59, "неуд", 1, 4);
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester)
VALUES (1, 2, 3, 59, "неуд", 0, 4);
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester)
VALUES (2, 2, 3, 59, "неуд", 0, 4);
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester)
VALUES (4, 2, 3, 59, "неуд", 0, 4);
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester)
VALUES (5, 2, 3, 59, "неуд", 0, 4);
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester)
VALUES (7, 2, 3, 59, "неуд", 0, 4);
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester)
VALUES (8, 2, 3, 59, "неуд", 0, 4);

INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester)
VALUES (10, 1, 6, 59, "неуд", 1, 4);
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester)
VALUES (11, 1, 6, 59, "неуд", 1, 4);
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester)
VALUES (13, 1, 6, 59, "неуд", 1, 4);
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester)
VALUES (14, 1, 6, 59, "неуд", 1, 4);
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester)
VALUES (10, 2, 6, 59, "неуд", 0, 4);
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester)
VALUES (11, 2, 6, 59, "неуд", 0, 4);
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester)
VALUES (13, 2, 6, 59, "неуд", 0, 4);
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester)
VALUES (14, 2, 6, 59, "неуд", 0, 4);