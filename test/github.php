<?php
foreach ($_SERVER as $key => $value) {
  file_put_contents('log.txt',$key.' => '.$value."\n",FILE_APPEND);
}
foreach ($_POST as $key => $value) {
  file_put_contents('log.txt',$key.' => '.$value."\n",FILE_APPEND);
}
?>
