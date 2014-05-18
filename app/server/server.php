<?php
$j = json_decode(file_get_contents('php://input'));

echo json_encode(array(
    'result' => array(
        'num' => array_reduce($j->parts, function($x, $y) {
            switch($y->operator) {
            case 'add': return $x + $y->num;
            case 'sub': return $x - $y->num;
            case 'multi': return $x * $y->num;
            case 'div': return $x / $y->num;
            case 'mod': return $x % $y->num;
            }
        }, 1)
    )
));