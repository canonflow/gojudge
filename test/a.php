<?php

$handle = fopen ("php://stdin","rb");
$line = fgets($handle);
echo "Input >> " . $line;
echo "\n";

for ($i = 0; $i < 10; $i++) {
    echo $i . "\n";
}
// if(trim($line) != 'yes'){
//     echo "ABORTING!\n";
// } else {
//     echo "\n";
//     echo "Thank you, continuing...\n";
//     echo "Input >> " . $line;
// }