<?php

$startTime = date('Y-m-d H:i:s', time());
$str = "Start from test_1.php at {$startTime}";
$file = fopen("log","a"); //開啟檔案
fwrite($file,$str);
fclose($file);
echo "{$str}\n";

sleep(30);

$endTime = date('Y-m-d H:i:s', time());
$str = "End from test_1.php at {$endTime}";
$file = fopen("log","a"); //開啟檔案
fwrite($file,$str);
fclose($file);
echo "{$str}\n";