CREATE TABLE IF NOT EXISTS reports (
    banner_id int,
    group_id int,
    slot_id int,
    FOREIGN KEY (banner_id) REFERENCES banners(id),
    FOREIGN KEY (group_id) REFERENCES groups(id),
    FOREIGN KEY (slot_id) REFERENCES slots(id),
    show int,
    conversion int,
    PRIMARY KEY(banner_id, group_id, slot_id)
);
