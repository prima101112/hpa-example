<?php

$dbh = new PDO('pgsql:host=db;dbname=example', 'example', 'example');

$sql = 'SELECT * FROM link';
foreach ($dbh->query($sql) as $row) {
    print $row['id'] . "\t|\t";
    print $row['url'] . "\t|\t";
    print $row['name'] . "\t|\t";
    echo "<hr>";
}

?>