CREATE TABLE IF NOT EXISTS rotations (
    banner_id int,
    slot_id int,
    FOREIGN KEY (banner_id) REFERENCES banners(id),
    FOREIGN KEY (slot_id) REFERENCES slots(id),
    title text,
    PRIMARY KEY(banner_id, slot_id)
);
