New-Item -ItemType Directory -Path . -Name "solutions"
New-Item -ItemType Directory -Path . -Name "inputs"

For ($i = 1; $i -le 12; $i++) {
    $tag = "day_"
    if ($i -lt 10) {
        $tag = $tag + "0"
    }

    New-Item -ItemType File -Path "./solutions" -Name ($tag + $i + ".go")
    New-Item -ItemType File -Path "./inputs" -Name  ($tag + $i + ".txt")
    New-Item -ItemType File -Path "./inputs" -Name ($tag + $i + ".test.txt")
}

PowerShell -NoExit