INSERT into vehicles(reg, type, owner, allowed) VALUES ('VIG 9619', 'Car','Bob Smith', 1);
INSERT into vehicles(reg, type, owner, allowed) VALUES ('BYZ 9480', 'Car','Alan Partridge', 1);
INSERT into vehicles(reg, type, owner, allowed) VALUES ('ABC 123', 'Car','Awesome Group Inc.', 1);
INSERT into vehicles(reg, type, owner, allowed) VALUES ('BF15 ZXY', 'Car','Important CEO Person', 1);

INSERT INTO events (time, type, reg) VALUES ('2023-08-02 08:51:00', 'Entry', 'VIG 9619');
INSERT INTO events (time, type, reg) VALUES ('2023-08-02 08:53:00', 'Entry', 'BYZ 9480');
INSERT INTO events (time, type, reg) VALUES ('2023-08-02 14:30:00', 'Entry', 'ABC 123');
INSERT INTO events (time, type, reg) VALUES ('2023-08-02 14:45:00', 'Exit', 'ABC 123');
INSERT INTO events (time, type, reg) VALUES ('2023-08-02 14:55:00', 'Entry', 'BF15 ZXY');
INSERT INTO events (time, type, reg) VALUES ('2023-08-02 15:00:00', 'Entry', 'AB12 CDE');
INSERT INTO events (time, type, reg) VALUES ('2023-08-02 15:00:00', 'Entry', 'DC66 CEA');
INSERT INTO events (time, type, reg) VALUES ('2023-08-02 15:22:00', 'Exit', 'BF15 ZXY');
INSERT INTO events (time, type, reg) VALUES ('2023-08-02 16:05:00', 'Exit', 'AB12 CDE');

INSERT INTO integrations (type, event_type, destination) VALUES ('email','Entry','hello@example.com');
INSERT INTO integrations (type, event_type, destination) VALUES ('email','Exit','hello@example.com');
