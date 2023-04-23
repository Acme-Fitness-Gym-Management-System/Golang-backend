-- name: GetClassEnrolmentByWeek :many
SELECT c.day, SUM(enrollment) as avg_enrollmen FROM class c
    JOIN (
    SELECT
    classcatalogue.courseid,
    COUNT(*) as enrollment
    FROM
    classcatalogue
    JOIN class ON classcatalogue.courseid = class.id
    GROUP BY
    classcatalogue.courseid
    ) AS enrollment ON enrollment.courseid = class.id
GROUP BY
    c.day
ORDER BY
    day ASC;


--name: GetClassEnrolmentTrend :many
SELECT
    date_trunc('week', enrolmentdate) AS week_start,
    COUNT(*) AS enrollment
FROM
    classcatalogue
WHERE
        enrolmentdate >= date_trunc('week', now() - interval '12 weeks')
GROUP BY
    week_start
ORDER BY
    week_start ASC;


--name GetVisitorsByHourOnCurrentDay :many
SELECT extract(hour from checkin) as hour, COUNT(*) FROM checkinactivity
WHERE date_trunc('day', checkin) = date_trunc('day', now())
GROUP BY hour;


--name GetVisitorsByHourOnCurrentWeek :many
SELECT extract(dow from checkin) as weekday, extract(hour from checkin) as hour, COUNT(*) FROM checkinactivity
WHERE checkin >= date_trunc('week', now()) AND checkin <= now()
GROUP BY weekday, hour;