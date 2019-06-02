DELETE
from marks;
ALTER TABLE marks
    AUTO_INCREMENT = 0;
DELETE
FROM loads;
ALTER TABLE loads
    AUTO_INCREMENT = 0;
DELETE
FROM discipline;
ALTER TABLE discipline
    AUTO_INCREMENT = 0;
# DELETE FROM sensitive_data;
# ALTER TABLE sensitive_data AUTO_INCREMENT=0;
DELETE
FROM student;
ALTER TABLE student
    AUTO_INCREMENT = 0;
DELETE
FROM employee;
ALTER TABLE employee
    AUTO_INCREMENT = 0;
DELETE
FROM accession;
ALTER TABLE accession
    AUTO_INCREMENT = 0;
DELETE
FROM people;
ALTER TABLE people
    AUTO_INCREMENT = 0;
DELETE
FROM speciality;
ALTER TABLE speciality
    AUTO_INCREMENT = 0;
DELETE
FROM direction;
ALTER TABLE direction
    AUTO_INCREMENT = 0;
DELETE
FROM cathedra;
ALTER TABLE cathedra
    AUTO_INCREMENT = 0;
DELETE
FROM faculty;
ALTER TABLE faculty
    AUTO_INCREMENT = 0;
DELETE
FROM institute;
ALTER TABLE institute
    AUTO_INCREMENT = 0;
-- DELETE FROM ranks;
-- ALTER TABLE ranks AUTO_INCREMENT=0;
DELETE
FROM groups;
ALTER TABLE groups
    AUTO_INCREMENT = 0;
DELETE
FROM loads_semester;
ALTER TABLE loads_semester
    AUTO_INCREMENT = 0;
-- DELETE FROM status;
-- ALTER TABLE status AUTO_INCREMENT=0;


INSERT INTO institute(name)
VALUES ("Навчально-науковий технологічний інститут харчової промисловості ім. М.В. Ломоносова");
INSERT INTO institute(name)
VALUES ("Навчально-науковий інститут комп`ютерних систем і технологій \"Індустрія 4.0\" ім. П.М. Платонова");
INSERT INTO institute(name)
VALUES ("Навчально-науковий інститут прикладної економіки та менеджменту ім. Г.Е. Вейнштейна");
INSERT INTO institute(name)
VALUES ("Навчально-науковий інститут холоду, кріотехнологій та екоенергетики ім. Мартиновського В.С.");


INSERT INTO faculty(name, id_institute)
VALUES ("Факультет технології зерна і зернового бізнесу", 1);
INSERT INTO faculty(name, id_institute)
VALUES ("Факультет інноваційних технологій харчування і ресторанно-готельного бізнесу", 1);
INSERT INTO faculty(name, id_institute)
VALUES ("Факультет технології та товарознавства харчових продуктів і продовольчого бізнесу", 1);
INSERT INTO faculty(name, id_institute)
VALUES ("Факультет технології вина та туристичного бізнесу", 1);
INSERT INTO faculty(name, id_institute)
VALUES ("Факультет комп`ютерних систем та автоматизації", 2);
INSERT INTO faculty(name, id_institute)
VALUES ("Факультет економіки, бізнесу і контролю", 3);
INSERT INTO faculty(name, id_institute)
VALUES ("Факультет менеджменту, маркетингу і логістики", 3);
INSERT INTO faculty(name, id_institute)
VALUES ("Факультет низькотемпературної техніки та інженерної механіки", 4);
INSERT INTO faculty(name, id_institute)
VALUES ("Факультет нафти, газу та екології", 4);
INSERT INTO faculty(name, id_institute)
VALUES ("Факультет комп`ютерної інженерії, програмування та кіберзахисту", 2);


INSERT INTO cathedra(name, id_faculty)
VALUES ("Кафедра технології переробки зерна", 1);
INSERT INTO cathedra(name, id_faculty)
VALUES ("Кафедра технології зберігання зерна", 1);
INSERT INTO cathedra(name, id_faculty)
VALUES ("Кафедра технологічного обладнання зернових виробництв", 1);
INSERT INTO cathedra(name, id_faculty)
VALUES ("Кафедра технології комбікормів і біопалива", 1);
INSERT INTO cathedra(name, id_faculty)
VALUES ("Кафедра управління бізнесом", 1);
INSERT INTO cathedra(name, id_faculty)
VALUES ("Кафедра технології хліба, кондитерських, макаронних виробів і харчоконцентратів", 1);
INSERT INTO cathedra(name, id_faculty)
VALUES ("Кафедра технології ресторанного і оздоровчого харчування", 2);
INSERT INTO cathedra(name, id_faculty)
VALUES ("Кафедра готельно-ресторанного бізнесу", 2);
INSERT INTO cathedra(name, id_faculty)
VALUES ("Кафедра інженерної графіки та технічного дизайну", 2);
INSERT INTO cathedra(name, id_faculty)
VALUES ("Кафедра безпеки життєдіяльності", 2);
INSERT INTO cathedra(name, id_faculty)
VALUES ("Кафедра іноземних мов", 2);
INSERT INTO cathedra(name, id_faculty)
VALUES ("Кафедра технології м´яса, риби і морепродуктів", 3);
INSERT INTO cathedra(name, id_faculty)
VALUES ("Кафедра технології молочних, олійно-жирових продуктів і косметики", 3);
INSERT INTO cathedra(name, id_faculty)
VALUES ("Кафедра харчової хімії та експертизи", 3);
INSERT INTO cathedra(name, id_faculty)
VALUES ("Кафедра економічної теорії та фінансово-економічної безпеки", 3);
INSERT INTO cathedra(name, id_faculty)
VALUES ("Кафедра товарознавства та митної справи", 3);
INSERT INTO cathedra(name, id_faculty)
VALUES ("Кафедра технології вина та енології", 4);
INSERT INTO cathedra(name, id_faculty)
VALUES ("Кафедра біохімії, мікробіології та фізіології харчування", 4);
INSERT INTO cathedra(name, id_faculty)
VALUES ("Кафедра біоінженерії і води", 4);
INSERT INTO cathedra(name, id_faculty)
VALUES ("Кафедра туристичного бізнесу та рекреації", 4);
INSERT INTO cathedra(name, id_faculty)
VALUES ("Кафедра автоматизації технологічних процесів і робототехнічних систем", 5);
INSERT INTO cathedra(name, id_faculty)
VALUES ("Кафедра електромеханіки та мехатроніки", 5);
INSERT INTO cathedra(name, id_faculty)
VALUES ("Кафедра вищої та прикладної математики", 5);
INSERT INTO cathedra(name, id_faculty)
VALUES ("Кафедра обліку і аудиту", 6);
INSERT INTO cathedra(name, id_faculty)
VALUES ("Кафедра економіки промисловості", 6);
INSERT INTO cathedra(name, id_faculty)
VALUES ("Кафедра українознавства і лінгводидактики", 6);
INSERT INTO cathedra(name, id_faculty)
VALUES ("Кафедра менеджменту і логістики", 7);
INSERT INTO cathedra(name, id_faculty)
VALUES ("Кафедра маркетингу, підприємництва і торгівлі", 7);
INSERT INTO cathedra(name, id_faculty)
VALUES ("Кафедра соціології, філософії і права", 7);
INSERT INTO cathedra(name, id_faculty)
VALUES ("Кафедра компресорів та пневмоагрегатів", 8);
INSERT INTO cathedra(name, id_faculty)
VALUES ("Кафедра кріогенної техніки", 8);
INSERT INTO cathedra(name, id_faculty)
VALUES ("Кафедра холодильних установок і кондиціювання повітря", 8);
INSERT INTO cathedra(name, id_faculty)
VALUES ("Кафедра процесів, обладнання та енергетичного менеджменту", 8);
INSERT INTO cathedra(name, id_faculty)
VALUES ("Кафедра фізики i матеріалознавства", 8);
INSERT INTO cathedra(name, id_faculty)
VALUES ("Кафедра термодинаміки та відновлювальної енергетики", 9);
INSERT INTO cathedra(name, id_faculty)
VALUES ("Кафедра теплоенергетики і трубопровідного транспорту енергоносіїв", 9);
INSERT INTO cathedra(name, id_faculty)
VALUES ("Кафедра екології та природоохоронні технології", 9);
INSERT INTO cathedra(name, id_faculty)
VALUES ("Кафедра теплофізики та прикладної екології", 9);
INSERT INTO cathedra(name, id_faculty)
VALUES ("Кафедра комп`ютерної інженерії", 10);
INSERT INTO cathedra(name, id_faculty)
VALUES ("Кафедра інформаційних технологій та кібербезпеки", 10);
INSERT INTO cathedra(name, id_faculty)
VALUES ("Кафедра фізичної культури та спорту", 10);

INSERT INTO direction(name, id_cathedra)
VALUES ("Комп''''''''ютерні науки", 40);
INSERT INTO direction(name, id_cathedra)
VALUES ("Комп''''''''ютерна інжеренерія", 40);

# INSERT INTO speciality(name, id_direction) VALUES ("speciality1", 1);
# INSERT INTO speciality(name, id_direction) VALUES ("speciality2", 1);
# INSERT INTO speciality(name, id_direction) VALUES ("speciality3", 2);
# INSERT INTO speciality(name, id_direction) VALUES ("speciality4", 2);
# INSERT INTO speciality(name, id_direction) VALUES ("speciality5", 3);
# INSERT INTO speciality(name, id_direction) VALUES ("speciality6", 3);
# INSERT INTO speciality(name, id_direction) VALUES ("speciality7", 4);
# INSERT INTO speciality(name, id_direction) VALUES ("speciality8", 4);
# INSERT INTO speciality(name, id_direction) VALUES ("speciality9", 5);
# INSERT INTO speciality(name, id_direction) VALUES ("speciality10", 5);
# INSERT INTO speciality(name, id_direction) VALUES ("speciality11", 6);
# INSERT INTO speciality(name, id_direction) VALUES ("speciality12", 6);
# INSERT INTO speciality(name, id_direction) VALUES ("speciality13", 7);
# INSERT INTO speciality(name, id_direction) VALUES ("speciality14", 7);
# INSERT INTO speciality(name, id_direction) VALUES ("speciality15", 8);
# INSERT INTO speciality(name, id_direction) VALUES ("speciality16", 8);
# INSERT INTO speciality(name, id_direction) VALUES ("speciality17", 9);
# INSERT INTO speciality(name, id_direction) VALUES ("speciality18", 9);
# INSERT INTO speciality(name, id_direction) VALUES ("speciality19", 10);
# INSERT INTO speciality(name, id_direction) VALUES ("speciality20", 10);
# INSERT INTO speciality(name, id_direction) VALUES ("speciality21", 11);
# INSERT INTO speciality(name, id_direction) VALUES ("speciality22", 11);
# INSERT INTO speciality(name, id_direction) VALUES ("speciality23", 12);
# INSERT INTO speciality(name, id_direction) VALUES ("speciality24", 12);
# INSERT INTO speciality(name, id_direction) VALUES ("speciality25", 13);
# INSERT INTO speciality(name, id_direction) VALUES ("speciality26", 13);
# INSERT INTO speciality(name, id_direction) VALUES ("speciality27", 14);
# INSERT INTO speciality(name, id_direction) VALUES ("speciality28", 14);
# INSERT INTO speciality(name, id_direction) VALUES ("speciality29", 15);
# INSERT INTO speciality(name, id_direction) VALUES ("speciality30", 15);
# INSERT INTO speciality(name, id_direction) VALUES ("speciality31", 16);
# INSERT INTO speciality(name, id_direction) VALUES ("speciality32", 16);

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

INSERT INTO people(fio, birthday, gender, img, comment, password, phone_number, email, sensitive_data, id_status,
                   have_access)
VALUES ("Степул Артем Мартиросович", "1997-06-12", 1, "/photo/stepul.jpg",
        "Співає у студ. клубі. Приймає актувну участь у заходах академії",
        "$2a$10$oqygBXUWPKkdL.wdac68n.GtjRGxk8LnaqxpTcNvZGWxRTtRHE0TK", "+380999999999", "artem.stepul@gmail.com",
        "JIW4cQpUP7Xcf+tB4C0HOrgJqc2EcC6w5UqnDugSvTEPgcmBEjTyNaQAjzjtBm2CGqjMUrjQ942X1h8vrTi/hRq3x4eu176e/0FMH5hJ12ljvX2IJcvs7AD1VL3zrenXuMhZLsdmWxqCBZRwxv2lZRJNrg6Uv8uymcVnhtcGVWy9+YQCQ87m+tufVtETec6Tdjk8C3HEs/DseyueKRzfk1CZI9+k6P88uCUhRM2XMmgMCj4Mm0EsO52nvHmk+/+fwiYt700qfknjVxG/SZAJvw5mPpHDEfG2wVJhBRhLHMBx23YxINd0JBPxMW6fT5yrYN+AqeI2amXcjAzLyGlC2Q==",
        1, 1);
INSERT INTO people(fio, birthday, gender, img, comment, password, phone_number, email, sensitive_data, id_status,
                   have_access)
VALUES ("Луценко Артем Геннадійович", "1997-03-16", 1, "/photo/default.png",
        "Займає посаду заступника голови студенської ради факультету. Приймає активну участь у суспільному житті академії",
        "$2a$10$oqygBXUWPKkdL.wdac68n.GtjRGxk8LnaqxpTcNvZGWxRTtRHE0TK", "+380333333333", "artem.lucenko@gmail.com",
        "JIW4cQpUP7Xcf+tB4C0HOrgJqc2EcC6w5UqnDugSvTEPgcmBEjTyNaQAjzjtBm2CGqjMUrjQ942X1h8vrTi/hRq3x4eu176e/0FMH5hJ12ljvX2IJcvs7AD1VL3zrenXuMhZLsdmWxqCBZRwxv2lZRJNrg6Uv8uymcVnhtcGVWy9+YQCQ87m+tufVtETec6Tdjk8C3HEs/DseyueKRzfk1CZI9+k6P88uCUhRM2XMmgMCj4Mm0EsO52nvHmk+/+fwiYt700qfknjVxG/SZAJvw5mPpHDEfG2wVJhBRhLHMBx23YxINd0JBPxMW6fT5yrYN+AqeI2amXcjAzLyGlC2Q==",
        1, 1);
INSERT INTO people(fio, birthday, gender, img, comment, password, phone_number, email, sensitive_data, id_status,
                   have_access)
VALUES ("Адміністратор", "1989-09-17", 1, "/photo/admin.png", "",
        "$2a$10$oqygBXUWPKkdL.wdac68n.GtjRGxk8LnaqxpTcNvZGWxRTtRHE0TK", "+380444444444", "admin@admin.com",
        "JIW4cQpUP7Xcf+tB4C0HOrgJqc2EcC6w5UqnDugSvTEPgcmBEjTyNaQAjzjtBm2CGqjMUrjQ942X1h8vrTi/hRq3x4eu176e/0FMH5hJ12ljvX2IJcvs7AD1VL3zrenXuMhZLsdmWxqCBZRwxv2lZRJNrg6Uv8uymcVnhtcGVWy9+YQCQ87m+tufVtETec6Tdjk8C3HEs/DseyueKRzfk1CZI9+k6P88uCUhRM2XMmgMCj4Mm0EsO52nvHmk+/+fwiYt700qfknjVxG/SZAJvw5mPpHDEfG2wVJhBRhLHMBx23YxINd0JBPxMW6fT5yrYN+AqeI2amXcjAzLyGlC2Q==",
        3, 1);
INSERT INTO people(fio, birthday, gender, img, comment, password, phone_number, email, sensitive_data, id_status,
                   have_access)
VALUES ("Єнов Богдан Олександрович", "1997-06-12", 1, "/photo/enov.jpg",
        "Відповідальний студент. Натхненно ставиться до навчання.",
        "$2a$10$oqygBXUWPKkdL.wdac68n.GtjRGxk8LnaqxpTcNvZGWxRTtRHE0TK", "+380999999999", "st3@gmail.com",
        "JIW4cQpUP7Xcf+tB4C0HOrgJqc2EcC6w5UqnDugSvTEPgcmBEjTyNaQAjzjtBm2CGqjMUrjQ942X1h8vrTi/hRq3x4eu176e/0FMH5hJ12ljvX2IJcvs7AD1VL3zrenXuMhZLsdmWxqCBZRwxv2lZRJNrg6Uv8uymcVnhtcGVWy9+YQCQ87m+tufVtETec6Tdjk8C3HEs/DseyueKRzfk1CZI9+k6P88uCUhRM2XMmgMCj4Mm0EsO52nvHmk+/+fwiYt700qfknjVxG/SZAJvw5mPpHDEfG2wVJhBRhLHMBx23YxINd0JBPxMW6fT5yrYN+AqeI2amXcjAzLyGlC2Q==",
        1, 1);
INSERT INTO people(fio, birthday, gender, img, comment, password, phone_number, email, sensitive_data, id_status,
                   have_access)
VALUES ("Вдовиченко Максим Ігорович", "1997-03-16", 1, "/photo/vdovichenko.jpg",
        "На відмінно проявив свої розумові здібності за час навчання.",
        "$2a$10$oqygBXUWPKkdL.wdac68n.GtjRGxk8LnaqxpTcNvZGWxRTtRHE0TK", "+380333333333", "vdovichenko@gmail.com",
        "JIW4cQpUP7Xcf+tB4C0HOrgJqc2EcC6w5UqnDugSvTEPgcmBEjTyNaQAjzjtBm2CGqjMUrjQ942X1h8vrTi/hRq3x4eu176e/0FMH5hJ12ljvX2IJcvs7AD1VL3zrenXuMhZLsdmWxqCBZRwxv2lZRJNrg6Uv8uymcVnhtcGVWy9+YQCQ87m+tufVtETec6Tdjk8C3HEs/DseyueKRzfk1CZI9+k6P88uCUhRM2XMmgMCj4Mm0EsO52nvHmk+/+fwiYt700qfknjVxG/SZAJvw5mPpHDEfG2wVJhBRhLHMBx23YxINd0JBPxMW6fT5yrYN+AqeI2amXcjAzLyGlC2Q==",
        1, 1);
INSERT INTO people(fio, birthday, gender, img, comment, password, phone_number, email, sensitive_data, id_status,
                   have_access)
VALUES ("Ольшевська Ольга Володимирівна", "1989-09-17", 0, "/photo/olshevska.jpg", "З 2017 року є секретар постійно діючої комісії з наукових досліджень при Вченій раді ОНАХТ.
З 2018 року член постійно діючої комісії з моніторингу поточних змін рейтингових показників при Вченій раді ОНАХТ.
Має понад 40 наукових публікацій та 11 навчально-методичних робіт, є постійним учасником та спікером міжнародних конференцій як на території України, так і за її межами.",
        "$2a$10$oqygBXUWPKkdL.wdac68n.GtjRGxk8LnaqxpTcNvZGWxRTtRHE0TK", "+380667634118", "olshevska.olga@gmail.com",
        "JIW4cQpUP7Xcf+tB4C0HOrgJqc2EcC6w5UqnDugSvTEPgcmBEjTyNaQAjzjtBm2CGqjMUrjQ942X1h8vrTi/hRq3x4eu176e/0FMH5hJ12ljvX2IJcvs7AD1VL3zrenXuMhZLsdmWxqCBZRwxv2lZRJNrg6Uv8uymcVnhtcGVWy9+YQCQ87m+tufVtETec6Tdjk8C3HEs/DseyueKRzfk1CZI9+k6P88uCUhRM2XMmgMCj4Mm0EsO52nvHmk+/+fwiYt700qfknjVxG/SZAJvw5mPpHDEfG2wVJhBRhLHMBx23YxINd0JBPxMW6fT5yrYN+AqeI2amXcjAzLyGlC2Q==",
        3, 1);
INSERT INTO people(fio, birthday, gender, img, comment, password, phone_number, email, sensitive_data, id_status,
                   have_access)
VALUES ("Іоргачов Євген Юрійович", "1998-03-01", 1, "/photo/iorgachev.jpg",
        "Відповідально ставиться до навчального процессу. Допомагає академії у стороніх заходах.",
        "$2a$10$oqygBXUWPKkdL.wdac68n.GtjRGxk8LnaqxpTcNvZGWxRTtRHE0TK", "+380999999999", "iorgachev@gmail.com",
        "JIW4cQpUP7Xcf+tB4C0HOrgJqc2EcC6w5UqnDugSvTEPgcmBEjTyNaQAjzjtBm2CGqjMUrjQ942X1h8vrTi/hRq3x4eu176e/0FMH5hJ12ljvX2IJcvs7AD1VL3zrenXuMhZLsdmWxqCBZRwxv2lZRJNrg6Uv8uymcVnhtcGVWy9+YQCQ87m+tufVtETec6Tdjk8C3HEs/DseyueKRzfk1CZI9+k6P88uCUhRM2XMmgMCj4Mm0EsO52nvHmk+/+fwiYt700qfknjVxG/SZAJvw5mPpHDEfG2wVJhBRhLHMBx23YxINd0JBPxMW6fT5yrYN+AqeI2amXcjAzLyGlC2Q==",
        1, 1);
INSERT INTO people(fio, birthday, gender, img, comment, password, phone_number, email, sensitive_data, id_status,
                   have_access)
VALUES ("Ткачук Станіслав Віталійович", "1997-03-16", 1, "/photo/tkachuck.jpg",
        "У 2018 році допомагав в організації заходів у честь все української олімпіади з інформатики",
        "$2a$10$oqygBXUWPKkdL.wdac68n.GtjRGxk8LnaqxpTcNvZGWxRTtRHE0TK", "+380333333333", "st6@gmail.com",
        "JIW4cQpUP7Xcf+tB4C0HOrgJqc2EcC6w5UqnDugSvTEPgcmBEjTyNaQAjzjtBm2CGqjMUrjQ942X1h8vrTi/hRq3x4eu176e/0FMH5hJ12ljvX2IJcvs7AD1VL3zrenXuMhZLsdmWxqCBZRwxv2lZRJNrg6Uv8uymcVnhtcGVWy9+YQCQ87m+tufVtETec6Tdjk8C3HEs/DseyueKRzfk1CZI9+k6P88uCUhRM2XMmgMCj4Mm0EsO52nvHmk+/+fwiYt700qfknjVxG/SZAJvw5mPpHDEfG2wVJhBRhLHMBx23YxINd0JBPxMW6fT5yrYN+AqeI2amXcjAzLyGlC2Q==",
        1, 1);
INSERT INTO people(fio, birthday, gender, img, comment, password, phone_number, email, sensitive_data, id_status,
                   have_access)
VALUES ("Бодюл Олена Станіславівна", "1977-07-23", 0, "/photo/bodul.jpeg", "Поєднує викладацьку діяльність та роботу редактора фахового видання – науково-технічного журналу «Холодильна техніка та технологія» з 2008 року. Працює у проблемній науково-дослідній лабораторії холодильної техніки з 2018 року.
Має 36 наукових публікацій та 12 навчально-методичних робіт",
        "$2a$10$oqygBXUWPKkdL.wdac68n.GtjRGxk8LnaqxpTcNvZGWxRTtRHE0TK", "+380487209144", "olbodiul@ukr.net",
        "JIW4cQpUP7Xcf+tB4C0HOrgJqc2EcC6w5UqnDugSvTEPgcmBEjTyNaQAjzjtBm2CGqjMUrjQ942X1h8vrTi/hRq3x4eu176e/0FMH5hJ12ljvX2IJcvs7AD1VL3zrenXuMhZLsdmWxqCBZRwxv2lZRJNrg6Uv8uymcVnhtcGVWy9+YQCQ87m+tufVtETec6Tdjk8C3HEs/DseyueKRzfk1CZI9+k6P88uCUhRM2XMmgMCj4Mm0EsO52nvHmk+/+fwiYt700qfknjVxG/SZAJvw5mPpHDEfG2wVJhBRhLHMBx23YxINd0JBPxMW6fT5yrYN+AqeI2amXcjAzLyGlC2Q==",
        3, 1);
INSERT INTO people(fio, birthday, gender, img, comment, password, phone_number, email, sensitive_data, id_status,
                   have_access)
VALUES ("Юрченко Боголеп Эразмович", "1970-01-01", "1", "/photo/default.png",
        "За час навчання виявив середні здібності. Пропусків уроків без поважних причин не має",
        "$2a$10$oqygBXUWPKkdL.wdac68n.GtjRGxk8LnaqxpTcNvZGWxRTtRHE0TK", "+380112223344", "st15@gmail.com",
        "JIW4cQpUP7Xcf+tB4C0HOrgJqc2EcC6w5UqnDugSvTEPgcmBEjTyNaQAjzjtBm2CGqjMUrjQ942X1h8vrTi/hRq3x4eu176e/0FMH5hJ12ljvX2IJcvs7AD1VL3zrenXuMhZLsdmWxqCBZRwxv2lZRJNrg6Uv8uymcVnhtcGVWy9+YQCQ87m+tufVtETec6Tdjk8C3HEs/DseyueKRzfk1CZI9+k6P88uCUhRM2XMmgMCj4Mm0EsO52nvHmk+/+fwiYt700qfknjVxG/SZAJvw5mPpHDEfG2wVJhBRhLHMBx23YxINd0JBPxMW6fT5yrYN+AqeI2amXcjAzLyGlC2Q==",
        7, 0);
INSERT INTO people(fio, birthday, gender, img, comment, password, phone_number, email, sensitive_data, id_status,
                   have_access)
VALUES ("Клокачева Лора Протасовна", "1970-01-01", "0", "/photo/default.png",
        "За час навчання виявив середні здібності. Пропусків уроків без поважних причин не має",
        "$2a$10$oqygBXUWPKkdL.wdac68n.GtjRGxk8LnaqxpTcNvZGWxRTtRHE0TK", "+380998887755", "st16@gmail.com",
        "JIW4cQpUP7Xcf+tB4C0HOrgJqc2EcC6w5UqnDugSvTEPgcmBEjTyNaQAjzjtBm2CGqjMUrjQ942X1h8vrTi/hRq3x4eu176e/0FMH5hJ12ljvX2IJcvs7AD1VL3zrenXuMhZLsdmWxqCBZRwxv2lZRJNrg6Uv8uymcVnhtcGVWy9+YQCQ87m+tufVtETec6Tdjk8C3HEs/DseyueKRzfk1CZI9+k6P88uCUhRM2XMmgMCj4Mm0EsO52nvHmk+/+fwiYt700qfknjVxG/SZAJvw5mPpHDEfG2wVJhBRhLHMBx23YxINd0JBPxMW6fT5yrYN+AqeI2amXcjAzLyGlC2Q==",
        7, 0);
INSERT INTO people(fio, birthday, gender, img, comment, password, phone_number, email, sensitive_data, id_status,
                   have_access)
VALUES ("Смирнова Катерина Василівна", "1983-12-25", 0, "/photo/smirnova.jpeg", "З 2014 року CIO проекту WAFWAF TECH
Має більш 20 наукових публікацій та більш 10 навчально-методичних робіт

Наукові інтереси: Інформаційна безпека, розробка комплексних систем захисту веб-додатків, розробка веб-додатків, машинне навчання

Стажування: У 2018 році пройшла стажування в Південноукраїнському національному педагогічному університеті ім. К. Д. Ушинського на кафедрі прикладної математики та інформатики

Нагороди: Нагороджена подякою проректора з науково-педагогічної та навчальної роботи ОНАХТ (2013), подякою проректора з науково-педагогічної роботи та міжнародних зв’язків ОНАХТ (2018).",
        "$2a$10$oqygBXUWPKkdL.wdac68n.GtjRGxk8LnaqxpTcNvZGWxRTtRHE0TK", "720 91 14", "smirnova.kathrin@gmail.com",
        "JIW4cQpUP7Xcf+tB4C0HOrgJqc2EcC6w5UqnDugSvTEPgcmBEjTyNaQAjzjtBm2CGqjMUrjQ942X1h8vrTi/hRq3x4eu176e/0FMH5hJ12ljvX2IJcvs7AD1VL3zrenXuMhZLsdmWxqCBZRwxv2lZRJNrg6Uv8uymcVnhtcGVWy9+YQCQ87m+tufVtETec6Tdjk8C3HEs/DseyueKRzfk1CZI9+k6P88uCUhRM2XMmgMCj4Mm0EsO52nvHmk+/+fwiYt700qfknjVxG/SZAJvw5mPpHDEfG2wVJhBRhLHMBx23YxINd0JBPxMW6fT5yrYN+AqeI2amXcjAzLyGlC2Q==",
        6, 1);
INSERT INTO people(fio, birthday, gender, img, comment, password, phone_number, email, sensitive_data, id_status,
                   have_access)
VALUES ("Попков Денис Миколайович", "1982-06-03", 1, "/photo/popkov.jpg", "Має більше 15 наукових публікацій та більше 20 навчально-методичних робіт

Наукові інтереси: генетичні алгоритми, системи автоматизації географічних досліджень, навчальні програмні продукти

Стажування: У 2016 р. пройшов стажування у .ОНУ ім.. Мечникова, за фахом «Новітні інформаційні технології та безпека комп’ютерних мереж».",
        "$2a$10$oqygBXUWPKkdL.wdac68n.GtjRGxk8LnaqxpTcNvZGWxRTtRHE0TK", "720 91 14", "popkovdn@ukr.net",
        "JIW4cQpUP7Xcf+tB4C0HOrgJqc2EcC6w5UqnDugSvTEPgcmBEjTyNaQAjzjtBm2CGqjMUrjQ942X1h8vrTi/hRq3x4eu176e/0FMH5hJ12ljvX2IJcvs7AD1VL3zrenXuMhZLsdmWxqCBZRwxv2lZRJNrg6Uv8uymcVnhtcGVWy9+YQCQ87m+tufVtETec6Tdjk8C3HEs/DseyueKRzfk1CZI9+k6P88uCUhRM2XMmgMCj4Mm0EsO52nvHmk+/+fwiYt700qfknjVxG/SZAJvw5mPpHDEfG2wVJhBRhLHMBx23YxINd0JBPxMW6fT5yrYN+AqeI2amXcjAzLyGlC2Q==",
        8, 1);
INSERT INTO people(fio, birthday, gender, img, comment, password, phone_number, email, sensitive_data, id_status,
                   have_access)
VALUES ("Становська Тетяна Павлівна", "1951-11-25", 0, "/photo/stanovska.jpeg", "Має 108 наукових публікацій та 23 навчально-методичні робіти.

Наукові інтереси: системи автоматизації проектування, нейронні мережі, новітні досягнення в області інформаційних технологій.

Стажування: У 2016 р отримала посвідчення про підвищення кваліфікації у Вищій школі педагогічної майстерності ОНАХТ. У 2018 р пройшла підвищення кваліфікації в ОНПУ. Загальний обсяг підвищення кваліфікації становить 3 кредити ECTS.",
        "$2a$10$oqygBXUWPKkdL.wdac68n.GtjRGxk8LnaqxpTcNvZGWxRTtRHE0TK", "720 91 18", "vmplotnik@gmail.com",
        "JIW4cQpUP7Xcf+tB4C0HOrgJqc2EcC6w5UqnDugSvTEPgcmBEjTyNaQAjzjtBm2CGqjMUrjQ942X1h8vrTi/hRq3x4eu176e/0FMH5hJ12ljvX2IJcvs7AD1VL3zrenXuMhZLsdmWxqCBZRwxv2lZRJNrg6Uv8uymcVnhtcGVWy9+YQCQ87m+tufVtETec6Tdjk8C3HEs/DseyueKRzfk1CZI9+k6P88uCUhRM2XMmgMCj4Mm0EsO52nvHmk+/+fwiYt700qfknjVxG/SZAJvw5mPpHDEfG2wVJhBRhLHMBx23YxINd0JBPxMW6fT5yrYN+AqeI2amXcjAzLyGlC2Q==",
        3, 1);

INSERT INTO people(fio, birthday, gender, img, comment, password, phone_number, email, sensitive_data, id_status,
                   have_access)
VALUES ("Гнатюк Анастасія Андріївна", "1998-05-25", 0, "/photo/july.png",
        "Студентка здатна, дисциплінована, відповідальна, працьовита, чуйна. У відносинах з одногрупниками товариська, комунікабельна, авторитет в групі високий. У відносинах з викладачами ввічлива, сумлінно ставиться до доручень класного керівника, сама проявляє ініціативу, завжди може запропонувати свою допомогу. Бере активну участь у громадському житті групи.",
        "$2a$10$oqygBXUWPKkdL.wdac68n.GtjRGxk8LnaqxpTcNvZGWxRTtRHE0TK", "+3801234567890", "gnatuk@gmail.com",
        "JIW4cQpUP7Xcf+tB4C0HOrgJqc2EcC6w5UqnDugSvTEPgcmBEjTyNaQAjzjtBm2CGqjMUrjQ942X1h8vrTi/hRq3x4eu176e/0FMH5hJ12ljvX2IJcvs7AD1VL3zrenXuMhZLsdmWxqCBZRwxv2lZRJNrg6Uv8uymcVnhtcGVWy9+YQCQ87m+tufVtETec6Tdjk8C3HEs/DseyueKRzfk1CZI9+k6P88uCUhRM2XMmgMCj4Mm0EsO52nvHmk+/+fwiYt700qfknjVxG/SZAJvw5mPpHDEfG2wVJhBRhLHMBx23YxINd0JBPxMW6fT5yrYN+AqeI2amXcjAzLyGlC2Q==",
        1, 1);
INSERT INTO people(fio, birthday, gender, img, comment, password, phone_number, email, sensitive_data, id_status,
                   have_access)
VALUES ("Снігур Тетяна Сергіївна", "1998-04-11", 0, "/photo/snigur.jpeg", "Наукові інтереси: генетичні алгоритми, нейронні мережі, спортивне програмування

Стажування: У 2014 році пройшла стажування в Одеському національному університеті імені І.І. Мечникова за фахом «Об’єктно-орієнтоване програмування» У 2018 році пройшла стажування в Університеті інформатики та прикладних знань(WSIU), м. Лодзь, Польща за фахом «Дослідження та сучасні тенденції розвитку ринку ІТ-персоналу»

Нагороди: Нагороджена дипломом 2 ступеню (2015р), Подякою декана КІПтаКЗ (2018р)",
        "$2a$10$oqygBXUWPKkdL.wdac68n.GtjRGxk8LnaqxpTcNvZGWxRTtRHE0TK", "720-91-44", "snigur.tatyana@ukr.net",
        "JIW4cQpUP7Xcf+tB4C0HOrgJqc2EcC6w5UqnDugSvTEPgcmBEjTyNaQAjzjtBm2CGqjMUrjQ942X1h8vrTi/hRq3x4eu176e/0FMH5hJ12ljvX2IJcvs7AD1VL3zrenXuMhZLsdmWxqCBZRwxv2lZRJNrg6Uv8uymcVnhtcGVWy9+YQCQ87m+tufVtETec6Tdjk8C3HEs/DseyueKRzfk1CZI9+k6P88uCUhRM2XMmgMCj4Mm0EsO52nvHmk+/+fwiYt700qfknjVxG/SZAJvw5mPpHDEfG2wVJhBRhLHMBx23YxINd0JBPxMW6fT5yrYN+AqeI2amXcjAzLyGlC2Q==",
        3, 1);

INSERT INTO people(fio, birthday, gender, img, comment, password, phone_number, email, sensitive_data, id_status,
                   have_access)
VALUES ("Курченко Владислав Іванович", "1978-09-25", 1, "/photo/kurchenko.png", "",
        "$2a$10$oqygBXUWPKkdL.wdac68n.GtjRGxk8LnaqxpTcNvZGWxRTtRHE0TK", "720-91-44", "kurchenko@gmail.com",
        "JIW4cQpUP7Xcf+tB4C0HOrgJqc2EcC6w5UqnDugSvTEPgcmBEjTyNaQAjzjtBm2CGqjMUrjQ942X1h8vrTi/hRq3x4eu176e/0FMH5hJ12ljvX2IJcvs7AD1VL3zrenXuMhZLsdmWxqCBZRwxv2lZRJNrg6Uv8uymcVnhtcGVWy9+YQCQ87m+tufVtETec6Tdjk8C3HEs/DseyueKRzfk1CZI9+k6P88uCUhRM2XMmgMCj4Mm0EsO52nvHmk+/+fwiYt700qfknjVxG/SZAJvw5mPpHDEfG2wVJhBRhLHMBx23YxINd0JBPxMW6fT5yrYN+AqeI2amXcjAzLyGlC2Q==",
        9, 0);

INSERT INTO people(fio, birthday, gender, img, comment, password, phone_number, email, sensitive_data, id_status,
                   have_access)
VALUES ("Чубенко Іван Артемович", "1978-03-01", 1, "/photo/chubenko.png", "",
        "$2a$10$oqygBXUWPKkdL.wdac68n.GtjRGxk8LnaqxpTcNvZGWxRTtRHE0TK", "720-91-44", "chubenko@gmail.com",
        "JIW4cQpUP7Xcf+tB4C0HOrgJqc2EcC6w5UqnDugSvTEPgcmBEjTyNaQAjzjtBm2CGqjMUrjQ942X1h8vrTi/hRq3x4eu176e/0FMH5hJ12ljvX2IJcvs7AD1VL3zrenXuMhZLsdmWxqCBZRwxv2lZRJNrg6Uv8uymcVnhtcGVWy9+YQCQ87m+tufVtETec6Tdjk8C3HEs/DseyueKRzfk1CZI9+k6P88uCUhRM2XMmgMCj4Mm0EsO52nvHmk+/+fwiYt700qfknjVxG/SZAJvw5mPpHDEfG2wVJhBRhLHMBx23YxINd0JBPxMW6fT5yrYN+AqeI2amXcjAzLyGlC2Q==",
        9, 0);

INSERT INTO people(fio, birthday, gender, img, comment, password, phone_number, email, sensitive_data, id_status,
                   have_access)
VALUES ("Мохавіна Елла Вадимівна", "1965-12-29", 0, "/photo/mahovina.png", "",
        "$2a$10$oqygBXUWPKkdL.wdac68n.GtjRGxk8LnaqxpTcNvZGWxRTtRHE0TK", "720-91-44", "mahovina@gmail.com",
        "JIW4cQpUP7Xcf+tB4C0HOrgJqc2EcC6w5UqnDugSvTEPgcmBEjTyNaQAjzjtBm2CGqjMUrjQ942X1h8vrTi/hRq3x4eu176e/0FMH5hJ12ljvX2IJcvs7AD1VL3zrenXuMhZLsdmWxqCBZRwxv2lZRJNrg6Uv8uymcVnhtcGVWy9+YQCQ87m+tufVtETec6Tdjk8C3HEs/DseyueKRzfk1CZI9+k6P88uCUhRM2XMmgMCj4Mm0EsO52nvHmk+/+fwiYt700qfknjVxG/SZAJvw5mPpHDEfG2wVJhBRhLHMBx23YxINd0JBPxMW6fT5yrYN+AqeI2amXcjAzLyGlC2Q==",
        9, 0);


INSERT INTO employee(id_people, date_invite, id_cathedra, id_rank)
VALUES (3, "2007-07-22", 21, 5);
INSERT INTO employee(id_people, date_invite, id_cathedra, id_rank)
VALUES (6, "2007-07-22", 40, 5);
INSERT INTO employee(id_people, date_invite, id_cathedra, id_rank)
VALUES (9, "2007-07-22", 40, 3);
INSERT INTO employee(id_people, date_invite, id_cathedra, id_rank)
VALUES (12, "2007-07-22", 40, 1);
INSERT INTO employee(id_people, date_invite, id_cathedra, id_rank)
VALUES (13, "2007-07-22", 40, 3);
INSERT INTO employee(id_people, date_invite, id_cathedra, id_rank)
VALUES (14, "2007-07-22", 40, 5);
INSERT INTO employee(id_people, date_invite, id_cathedra, id_rank)
VALUES (16, "2007-07-22", 40, 1);
INSERT INTO employee(id_people, date_invite, id_cathedra, id_rank)
VALUES (17, "2001-07-22", 40, 1);
INSERT INTO employee(id_people, date_invite, id_cathedra, id_rank)
VALUES (18, "2003-07-22", 40, 1);
INSERT INTO employee(id_people, date_invite, id_cathedra, id_rank)
VALUES (19, "2002-07-22", 40, 1);

INSERT INTO groups(id_employee, name, id_direction)
VALUES (6, "311а", 1);
INSERT INTO groups(id_employee, name, id_direction)
VALUES (NULL, "311б", 1);
INSERT INTO groups(id_employee, name, id_direction)
VALUES (NULL, "312а", 1);
INSERT INTO groups(id_employee, name, id_direction)
VALUES (NULL, "312б", 1);
INSERT INTO groups(id_employee, name, id_direction)
VALUES (NULL, "321а", 1);
INSERT INTO groups(id_employee, name, id_direction)
VALUES (NULL, "321б", 1);
INSERT INTO groups(id_employee, name, id_direction)
VALUES (NULL, "322а", 1);
INSERT INTO groups(id_employee, name, id_direction)
VALUES (NULL, "322б", 1);
INSERT INTO groups(id_employee, name, id_direction)
VALUES (NULL, "331а", 1);
INSERT INTO groups(id_employee, name, id_direction)
VALUES (NULL, "331б", 1);
INSERT INTO groups(id_employee, name, id_direction)
VALUES (NULL, "332а", 1);
INSERT INTO groups(id_employee, name, id_direction)
VALUES (NULL, "332б", 1);
INSERT INTO groups(id_employee, name, id_direction)
VALUES (NULL, "333а", 1);
INSERT INTO groups(id_employee, name, id_direction)
VALUES (NULL, "333б", 1);
INSERT INTO groups(id_employee, name, id_direction)
VALUES (NULL, "334а", 1);
INSERT INTO groups(id_employee, name, id_direction)
VALUES (NULL, "334б", 1);
INSERT INTO groups(id_employee, name, id_direction)
VALUES (NULL, "341а", 1);
INSERT INTO groups(id_employee, name, id_direction)
VALUES (NULL, "341б", 1);
INSERT INTO groups(id_employee, name, id_direction)
VALUES (3, "342а", 1);
INSERT INTO groups(id_employee, name, id_direction)
VALUES (NULL, "342б", 1);
INSERT INTO groups(id_employee, name, id_direction)
VALUES (NULL, "343а", 1);
INSERT INTO groups(id_employee, name, id_direction)
VALUES (6, "343б", 1);
INSERT INTO groups(id_employee, name, id_direction)
VALUES (NULL, "344а", 1);
INSERT INTO groups(id_employee, name, id_direction)
VALUES (NULL, "344б", 1);

INSERT INTO student(id_people, date_admission, is_full_time, is_cut, semester, id_group)
VALUES (1, "2017-09-01", 1, 1, 4, 22);
INSERT INTO student(id_people, date_admission, is_full_time, is_cut, semester, id_group)
VALUES (2, "2017-09-01", 1, 1, 4, 22);
INSERT INTO student(id_people, date_admission, is_full_time, is_cut, semester, id_group)
VALUES (4, "2017-09-01", 1, 1, 4, 22);
INSERT INTO student(id_people, date_admission, is_full_time, is_cut, semester, id_group)
VALUES (5, "2017-09-01", 1, 1, 4, 22);
INSERT INTO student(id_people, date_admission, is_full_time, is_cut, semester, id_group)
VALUES (7, "2017-09-01", 1, 1, 4, 22);
INSERT INTO student(id_people, date_admission, is_full_time, is_cut, semester, id_group)
VALUES (8, "2017-09-01", 1, 1, 4, 22);
INSERT INTO student(id_people, date_admission, is_full_time, is_cut, semester, id_group)
VALUES (10, "2014-09-01", 1, 1, 4, NULL);
INSERT INTO student(id_people, date_admission, is_full_time, is_cut, semester, id_group)
VALUES (11, "2014-09-01", 1, 1, 4, NULL);
INSERT INTO student(id_people, date_admission, is_full_time, is_cut, semester, id_group)
VALUES (15, "2017-09-01", 1, 1, 4, 22);

INSERT INTO accession(id_people, edit_access, set_absence, get_absence, set_mark, set_event, get_sensitive,
                      set_sensitive, manage_load, manage_academ)
VALUES (1, 0, 0, 0, 0, 0, 0, 0, 0, 0);
INSERT INTO accession(id_people, edit_access, set_absence, get_absence, set_mark, set_event, get_sensitive,
                      set_sensitive, manage_load, manage_academ)
VALUES (2, 0, 0, 0, 0, 0, 0, 0, 0, 0);
INSERT INTO accession(id_people, edit_access, set_absence, get_absence, set_mark, set_event, get_sensitive,
                      set_sensitive, manage_load, manage_academ)
VALUES (3, 1, 0, 0, 0, 0, 1, 1, 1, 1);
INSERT INTO accession(id_people, edit_access, set_absence, get_absence, set_mark, set_event, get_sensitive,
                      set_sensitive, manage_load, manage_academ)
VALUES (4, 0, 0, 0, 0, 0, 0, 0, 0, 0);
INSERT INTO accession(id_people, edit_access, set_absence, get_absence, set_mark, set_event, get_sensitive,
                      set_sensitive, manage_load, manage_academ)
VALUES (5, 0, 0, 0, 0, 0, 0, 0, 0, 0);
INSERT INTO accession(id_people, edit_access, set_absence, get_absence, set_mark, set_event, get_sensitive,
                      set_sensitive, manage_load, manage_academ)
VALUES (13, 0, 0, 0, 0, 0, 1, 1, 1, 0);
INSERT INTO accession(id_people, edit_access, set_absence, get_absence, set_mark, set_event, get_sensitive,
                      set_sensitive, manage_load, manage_academ)
VALUES (7, 0, 0, 0, 0, 0, 0, 0, 0, 0);
INSERT INTO accession(id_people, edit_access, set_absence, get_absence, set_mark, set_event, get_sensitive,
                      set_sensitive, manage_load, manage_academ)
VALUES (8, 0, 0, 0, 0, 0, 0, 0, 0, 0);
INSERT INTO accession(id_people, edit_access, set_absence, get_absence, set_mark, set_event, get_sensitive,
                      set_sensitive, manage_load, manage_academ)
VALUES (9, 0, 0, 0, 1, 0, 0, 0, 0, 0);
INSERT INTO accession(id_people, edit_access, set_absence, get_absence, set_mark, set_event, get_sensitive,
                      set_sensitive, manage_load, manage_academ)
VALUES (10, 0, 0, 0, 0, 0, 0, 0, 0, 0);
INSERT INTO accession(id_people, edit_access, set_absence, get_absence, set_mark, set_event, get_sensitive,
                      set_sensitive, manage_load, manage_academ)
VALUES (11, 0, 0, 0, 0, 0, 0, 0, 0, 0);
INSERT INTO accession(id_people, edit_access, set_absence, get_absence, set_mark, set_event, get_sensitive,
                      set_sensitive, manage_load, manage_academ)
VALUES (12, 0, 0, 1, 1, 1, 1, 0, 0, 0);
INSERT INTO accession(id_people, edit_access, set_absence, get_absence, set_mark, set_event, get_sensitive,
                      set_sensitive, manage_load, manage_academ)
VALUES (6, 0, 0, 0, 1, 0, 1, 0, 0, 0);
INSERT INTO accession(id_people, edit_access, set_absence, get_absence, set_mark, set_event, get_sensitive,
                      set_sensitive, manage_load, manage_academ)
VALUES (14, 0, 0, 0, 1, 0, 1, 0, 0, 0);
INSERT INTO accession(id_people, edit_access, set_absence, get_absence, set_mark, set_event, get_sensitive,
                      set_sensitive, manage_load, manage_academ)
VALUES (15, 0, 0, 0, 0, 0, 0, 0, 0, 0);
INSERT INTO accession(id_people, edit_access, set_absence, get_absence, set_mark, set_event, get_sensitive,
                      set_sensitive, manage_load, manage_academ)
VALUES (16, 0, 0, 0, 1, 0, 0, 0, 0, 0);
INSERT INTO accession(id_people, edit_access, set_absence, get_absence, set_mark, set_event, get_sensitive,
                      set_sensitive, manage_load, manage_academ)
VALUES (17, 0, 0, 0, 0, 0, 0, 0, 0, 0);
INSERT INTO accession(id_people, edit_access, set_absence, get_absence, set_mark, set_event, get_sensitive,
                      set_sensitive, manage_load, manage_academ)
VALUES (18, 0, 0, 0, 0, 0, 0, 0, 0, 0);
INSERT INTO accession(id_people, edit_access, set_absence, get_absence, set_mark, set_event, get_sensitive,
                      set_sensitive, manage_load, manage_academ)
VALUES (19, 0, 0, 0, 0, 0, 0, 0, 0, 0);



# INSERT INTO sensitive_data(id_people, passport_code, rntrs, reg_address, military_id)
# VALUES (1, "122fasf2", "214124124", "ул. Пушкина, дом Колатушкина, квартира 25", "ком214124");
# INSERT INTO sensitive_data(id_people, passport_code, rntrs, reg_address, military_id)
# VALUES (2, "214124124", "122fasf2", "ул. Пушкина, дом Колатушкина, квартира 12", "ыфафаы251512");
# INSERT INTO sensitive_data(id_people, passport_code, rntrs, reg_address, military_id)
# VALUES (3, "цйкйкf2", "2141фафыафыа", "ул. Пушкина, дом Колатушкина, квартира 36", "пддждхзз215251аыф");
# INSERT INTO sensitive_data(id_people, passport_code, rntrs, reg_address, military_id)
# VALUES (4, "122fasf2", "214124124", "ул. Пушкина, дом Колатушкина, квартира 25", "ком214124");
# INSERT INTO sensitive_data(id_people, passport_code, rntrs, reg_address, military_id)
# VALUES (5, "214124124", "122fasf2", "ул. Пушкина, дом Колатушкина, квартира 12", "ыфафаы251512");
# INSERT INTO sensitive_data(id_people, passport_code, rntrs, reg_address, military_id)
# VALUES (6, "цйкйкf2", "2141фафыафыа", "ул. Пушкина, дом Колатушкина, квартира 36", "пддждхзз215251аыф");
# INSERT INTO sensitive_data(id_people, passport_code, rntrs, reg_address, military_id)
# VALUES (7, "122fasf2", "214124124", "ул. Пушкина, дом Колатушкина, квартира 25", "ком214124");
# INSERT INTO sensitive_data(id_people, passport_code, rntrs, reg_address, military_id)
# VALUES (8, "214124124", "122fasf2", "ул. Пушкина, дом Колатушкина, квартира 12", "ыфафаы251512");
# INSERT INTO sensitive_data(id_people, passport_code, rntrs, reg_address, military_id)
# VALUES (9, "цйкйкf2", "2141фафыафыа", "ул. Пушкина, дом Колатушкина, квартира 36", "пддждхзз215251аыф");
# INSERT INTO sensitive_data(id_people, passport_code, rntrs, reg_address, military_id)
# VALUES (10, "122fasf2", "214124124", "ул. Пушкина, дом Колатушкина, квартира 25", "ком214124");
# INSERT INTO sensitive_data(id_people, passport_code, rntrs, reg_address, military_id)
# VALUES (11, "214124124", "122fasf2", "ул. Пушкина, дом Колатушкина, квартира 12", "ыфафаы251512");
# INSERT INTO sensitive_data(id_people, passport_code, rntrs, reg_address, military_id)
# VALUES (12, "цйкйкf2", "2141фафыафыа", "ул. Пушкина, дом Колатушкина, квартира 36", "пддждхзз215251аыф");
# INSERT INTO sensitive_data(id_people, passport_code, rntrs, reg_address, military_id)
# VALUES (13, "122fasf2", "214124124", "ул. Пушкина, дом Колатушкина, квартира 25", "ком214124");
# INSERT INTO sensitive_data(id_people, passport_code, rntrs, reg_address, military_id)
# VALUES (14, "214124124", "122fasf2", "ул. Пушкина, дом Колатушкина, квартира 12", "ыфафаы251512");
# INSERT INTO sensitive_data(id_people, passport_code, rntrs, reg_address, military_id)
# VALUES (15, "фыафыафыа", "122fаыфафыаasf2", "ул. фыафыа, дом Колатушафыафыакина, кварафыафыатира 12", "ыфафааыфафыаы251512");
# INSERT INTO sensitive_data(id_people, passport_code, rntrs, reg_address, military_id)
# VALUES (16, "122fasf2", "214124124", "ул. Пушкина, дом Колатушкина, квартира 25", "ком214124");
# INSERT INTO sensitive_data(id_people, passport_code, rntrs, reg_address, military_id)
# VALUES (17, "214124124", "122fasf2", "ул. Пушкина, дом Колатушкина, квартира 12", "ыфафаы251512");

INSERT INTO discipline(name)
VALUES ("ММДО");
INSERT INTO discipline(name)
VALUES ("ОАПСОС");
INSERT INTO discipline(name)
VALUES ("КПП");
INSERT INTO discipline(name)
VALUES ("ССУБД");
INSERT INTO discipline(name)
VALUES ("САПКИС");
INSERT INTO discipline(name)
VALUES ("ТКП");
INSERT INTO discipline(name)
VALUES ("МтаСШИ");
INSERT INTO discipline(name)
VALUES ("ИАД");
INSERT INTO discipline(name)
VALUES ("ТЗИ");
INSERT INTO discipline(name)
VALUES ("ВДтаВТ");
INSERT INTO discipline(name)
VALUES ("ОБДЗ");
INSERT INTO discipline(name)
VALUES ("ТРСПО");
INSERT INTO discipline(name)
VALUES ("ЕС");

INSERT INTO loads_semester(start, end, name)
VALUES (1546263600, 1546293600, "1 семестр 2017-2018 року");
INSERT INTO loads_semester(start, end, name)
VALUES (1546263600, 1546293600, "2 семестр 2017-2018 року");
INSERT INTO loads_semester(start, end, name)
VALUES (1546263600, 1546293600, "1 семестр 2018-2019 року");
INSERT INTO loads_semester(start, end, name)
VALUES (1546293600, 1562446800, "2 семестр 2018-2019 року");

INSERT INTO loads(id_discipline, id_employee, id_group, semester, id_semester, id_assistant)
VALUES (1, 9, 22, 1, 1, 16);
INSERT INTO loads(id_discipline, id_employee, id_group, semester, id_semester, id_assistant)
VALUES (2, 9, 22, 1, 1, 16);
INSERT INTO loads(id_discipline, id_employee, id_group, semester, id_semester, id_assistant)
VALUES (3, 6, 22, 1, 1, NULL);
INSERT INTO loads(id_discipline, id_employee, id_group, semester, id_semester, id_assistant)
VALUES (11, 13, 22, 1, 1, NULL);
INSERT INTO loads(id_discipline, id_employee, id_group, semester, id_semester, id_assistant)
VALUES (6, 14, 22, 1, 1, NULL);

INSERT INTO loads(id_discipline, id_employee, id_group, semester, id_semester, id_assistant)
VALUES (1, 9, 22, 2, 2, 16);
INSERT INTO loads(id_discipline, id_employee, id_group, semester, id_semester, id_assistant)
VALUES (4, 13, 22, 2, 2, NULL);
INSERT INTO loads(id_discipline, id_employee, id_group, semester, id_semester, id_assistant)
VALUES (5, 13, 22, 2, 2, 16);

INSERT INTO loads(id_discipline, id_employee, id_group, semester, id_semester, id_assistant)
VALUES (6, 6, 22, 3, 3, 9);
INSERT INTO loads(id_discipline, id_employee, id_group, semester, id_semester, id_assistant)
VALUES (7, 6, 22, 3, 3, 9);
INSERT INTO loads(id_discipline, id_employee, id_group, semester, id_semester, id_assistant)
VALUES (12, 9, 22, 3, 3, 16);
INSERT INTO loads(id_discipline, id_employee, id_group, semester, id_semester, id_assistant)
VALUES (13, 12, 22, 3, 3, 16);

INSERT INTO loads(id_discipline, id_employee, id_group, semester, id_semester, id_assistant)
VALUES (8, 6, 22, 4, 4, 9);
INSERT INTO loads(id_discipline, id_employee, id_group, semester, id_semester, id_assistant)
VALUES (9, 16, 22, 4, 4, 9);
INSERT INTO loads(id_discipline, id_employee, id_group, semester, id_semester, id_assistant)
VALUES (10, 6, 22, 4, 4, 9);

INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (15, 1, 9, 75, "задовільно", 0, 1, "2017-12-05");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (15, 2, 9, 89, "добре", 1, 1, "2017-12-07");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (15, 3, 6, 64, "задовільно", 0, 1, "2017-12-09");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (15, 11, 13, 64, "задовільно", 1, 1, "2017-12-13");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (15, 6, 14, 64, "задовільно", 1, 1, "2017-12-15");

INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (15, 1, 13, 90, "добре", 1, 2, "2018-06-07");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (15, 4, 13, 83, "добре", 0, 2, "2018-06-05");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (15, 5, 13, 86, "добре", 1, 2, "2018-06-12");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (15, 3, 6, 64, "задовільно", 1, 2, "2018-06-09");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (15, 3, 6, 64, "задовільно", 0, 2, "2017-12-09");

INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (15, 6, 6, 75, "задовільно", 0, 3, "2018-12-05");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (15, 7, 6, 60, "задовільно", 1, 3, "2018-12-07");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (15, 12, 9, 75, "задовільно", 0, 3, "2018-12-09");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (15, 13, 12, 60, "задовільно", 1, 3, "2018-12-12");


INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (1, 1, 9, 75, "задовільно", 0, 1, "2017-12-05");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (1, 2, 9, 89, "добре", 1, 1, "2017-12-07");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (1, 3, 6, 64, "задовільно", 0, 1, "2017-12-09");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (1, 11, 13, 64, "задовільно", 1, 1, "2017-12-13");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (1, 6, 14, 64, "задовільно", 1, 1, "2017-12-15");


INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (1, 1, 13, 90, "добре", 1, 2, "2018-06-07");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (1, 4, 13, 83, "добре", 0, 2, "2018-06-05");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (1, 5, 13, 86, "добре", 1, 2, "2018-06-12");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (1, 3, 6, 64, "задовільно", 1, 2, "2018-06-09");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (1, 3, 6, 64, "задовільно", 0, 2, "2017-12-09");

INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (1, 6, 6, 75, "задовільно", 0, 3, "2018-12-05");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (1, 7, 6, 60, "задовільно", 1, 3, "2018-12-07");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (1, 12, 9, 75, "задовільно", 0, 3, "2018-12-09");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (1, 13, 12, 60, "задовільно", 1, 3, "2018-12-12");



INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (2, 1, 9, 75, "задовільно", 0, 1, "2017-12-05");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (2, 2, 9, 89, "добре", 1, 1, "2017-12-07");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (2, 3, 6, 64, "задовільно", 0, 1, "2017-12-09");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (2, 11, 13, 64, "задовільно", 1, 1, "2017-12-13");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (2, 6, 14, 64, "задовільно", 1, 1, "2017-12-15");


INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (2, 1, 13, 90, "добре", 1, 2, "2018-06-07");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (2, 4, 13, 83, "добре", 0, 2, "2018-06-05");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (2, 5, 13, 86, "добре", 1, 2, "2018-06-12");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (2, 3, 6, 64, "задовільно", 1, 2, "2018-06-09");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (2, 3, 6, 64, "задовільно", 0, 2, "2017-12-09");


INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (2, 6, 6, 75, "задовільно", 0, 3, "2018-12-05");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (2, 7, 6, 60, "задовільно", 1, 3, "2018-12-07");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (2, 12, 9, 75, "задовільно", 0, 3, "2018-12-09");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (2, 13, 12, 60, "задовільно", 1, 3, "2018-12-12");


INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (4, 1, 9, 75, "задовільно", 0, 1, "2017-12-05");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (4, 2, 9, 89, "добре", 1, 1, "2017-12-07");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (4, 3, 6, 64, "задовільно", 0, 1, "2017-12-09");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (4, 11, 13, 64, "задовільно", 1, 1, "2017-12-13");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (4, 6, 14, 64, "задовільно", 1, 1, "2017-12-15");


INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (4, 1, 13, 90, "добре", 1, 2, "2018-06-07");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (4, 4, 13, 83, "добре", 0, 2, "2018-06-05");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (4, 5, 13, 86, "добре", 1, 2, "2018-06-12");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (4, 3, 6, 64, "задовільно", 1, 2, "2018-06-09");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (4, 3, 6, 64, "задовільно", 0, 2, "2017-12-09");



INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (4, 6, 6, 75, "задовільно", 0, 3, "2018-12-05");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (4, 7, 6, 60, "задовільно", 1, 3, "2018-12-07");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (4, 12, 9, 75, "задовільно", 0, 3, "2018-12-09");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (4, 13, 12, 60, "задовільно", 1, 3, "2018-12-12");



INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (5, 1, 9, 75, "задовільно", 0, 1, "2017-12-05");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (5, 2, 9, 89, "добре", 1, 1, "2017-12-07");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (5, 3, 6, 64, "задовільно", 0, 1, "2017-12-09");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (5, 11, 13, 64, "задовільно", 1, 1, "2017-12-13");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (5, 6, 14, 64, "задовільно", 1, 1, "2017-12-15");


INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (5, 1, 13, 90, "добре", 1, 2, "2018-06-07");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (5, 4, 13, 83, "добре", 0, 2, "2018-06-05");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (5, 5, 13, 86, "добре", 1, 2, "2018-06-12");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (5, 3, 6, 64, "задовільно", 1, 2, "2018-06-09");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (5, 3, 6, 64, "задовільно", 0, 2, "2017-12-09");


INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (5, 6, 6, 75, "задовільно", 0, 3, "2018-12-05");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (5, 7, 6, 60, "задовільно", 1, 3, "2018-12-07");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (5, 12, 9, 75, "задовільно", 0, 3, "2018-12-09");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (5, 13, 12, 60, "задовільно", 1, 3, "2018-12-12");



INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (7, 1, 9, 75, "задовільно", 0, 1, "2017-12-05");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (7, 2, 9, 89, "добре", 1, 1, "2017-12-07");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (7, 3, 6, 64, "задовільно", 0, 1, "2017-12-09");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (7, 11, 13, 64, "задовільно", 1, 1, "2017-12-13");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (7, 6, 14, 64, "задовільно", 1, 1, "2017-12-15");


INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (7, 1, 13, 90, "добре", 1, 2, "2018-06-07");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (7, 4, 13, 83, "добре", 0, 2, "2018-06-05");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (7, 5, 13, 86, "добре", 1, 2, "2018-06-12");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (7, 3, 6, 64, "задовільно", 1, 2, "2018-06-09");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (7, 3, 6, 64, "задовільно", 0, 2, "2017-12-09");



INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (7, 6, 6, 75, "задовільно", 0, 3, "2018-12-05");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (7, 7, 6, 60, "задовільно", 1, 3, "2018-12-07");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (7, 12, 9, 75, "задовільно", 0, 3, "2018-12-09");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (7, 13, 12, 60, "задовільно", 1, 3, "2018-12-12");


INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (8, 1, 9, 75, "задовільно", 0, 1, "2017-12-05");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (8, 2, 9, 89, "добре", 1, 1, "2017-12-07");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (8, 3, 6, 64, "задовільно", 0, 1, "2017-12-09");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (8, 11, 13, 64, "задовільно", 1, 1, "2017-12-13");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (8, 6, 14, 64, "задовільно", 1, 1, "2017-12-15");


INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (8, 1, 13, 90, "добре", 1, 2, "2018-06-07");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (8, 4, 13, 83, "добре", 0, 2, "2018-06-05");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (8, 5, 13, 86, "добре", 1, 2, "2018-06-12");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (8, 3, 6, 64, "задовільно", 1, 2, "2018-06-09");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (8, 3, 6, 64, "задовільно", 0, 2, "2017-12-09");



INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (8, 6, 6, 75, "задовільно", 0, 3, "2018-12-05");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (8, 7, 6, 60, "задовільно", 1, 3, "2018-12-07");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (8, 12, 9, 75, "задовільно", 0, 3, "2018-12-09");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (8, 13, 12, 60, "задовільно", 1, 3, "2018-12-12");



INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (10, 1, 9, 75, "задовільно", 0, 1, "2017-12-05");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (10, 2, 9, 89, "добре", 1, 1, "2017-12-07");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (10, 3, 6, 64, "задовільно", 0, 1, "2017-12-09");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (10, 11, 13, 64, "задовільно", 1, 1, "2017-12-13");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (10, 6, 14, 64, "задовільно", 1, 1, "2017-12-15");


INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (10, 1, 13, 90, "добре", 1, 2, "2018-06-07");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (10, 4, 13, 83, "добре", 0, 2, "2018-06-05");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (10, 5, 13, 86, "добре", 1, 2, "2018-06-12");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (10, 3, 6, 64, "задовільно", 1, 1, "2018-06-09");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (10, 3, 6, 64, "задовільно", 0, 2, "2017-12-09");


INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (10, 6, 6, 75, "задовільно", 0, 3, "2018-12-05");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (10, 7, 6, 60, "задовільно", 1, 3, "2018-12-07");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (10, 12, 9, 75, "задовільно", 0, 3, "2018-12-09");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (10, 13, 12, 60, "задовільно", 1, 3, "2018-12-12");



INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (11, 1, 9, 75, "задовільно", 0, 1, "2017-12-05");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (11, 2, 9, 89, "добре", 1, 1, "2017-12-07");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (11, 3, 6, 64, "задовільно", 0, 1, "2017-12-09");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (11, 11, 13, 64, "задовільно", 1, 1, "2017-12-13");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (11, 6, 14, 64, "задовільно", 1, 1, "2017-12-15");


INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (11, 1, 13, 90, "добре", 1, 2, "2018-06-07");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (11, 4, 13, 83, "добре", 0, 2, "2018-06-05");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (11, 5, 13, 86, "добре", 1, 2, "2018-06-12");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (11, 3, 6, 64, "задовільно", 1, 1, "2018-06-09");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (11, 3, 6, 64, "задовільно", 0, 2, "2017-12-09");



INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (11, 6, 6, 75, "задовільно", 0, 3, "2018-12-05");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (11, 7, 6, 60, "задовільно", 1, 3, "2018-12-07");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (11, 12, 9, 75, "задовільно", 0, 3, "2018-12-09");
INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
VALUES (11, 13, 12, 60, "задовільно", 1, 3, "2018-12-12");


# INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
# VALUES (15, 8, 6, 97, "відмінно", 1, 4, "2018-09-05");
# INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
# VALUES (15, 9, 16, 78, "задовільно", 0, 4, "2018-09-05");
# INSERT INTO marks(id_student, id_discipline, id_employee, value, national_value, is_exam, semester, date)
# VALUES (15, 10, 6, 87, "добре", 1, 4, "2018-09-05");