<?php

$startTime = date('Y-m-d H:i:s', time());
$str = "Start from test_2.php at {$startTime}";
$file = fopen("log","a"); //開啟檔案
fwrite($file,$str);
fclose($file);
echo "{$str}\n";

sleep(5);

$endTime = date('Y-m-d H:i:s', time());
$str = "End from test_2.php at {$endTime}";
$file = fopen("log","a"); //開啟檔案
fwrite($file,$str);
fclose($file);
echo "{$str}\n";