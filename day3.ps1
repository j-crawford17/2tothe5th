# variables
$nums = @(2, 7, 11, 15)
$target = 18

function Get-twoSums($nums, $target){
    # Iterate through nums array
    for($i=0; $i -lt $nums.Length; $i++){
        # Iterate through nums array + 1
        for($j=$i+1; $j -lt $nums.Length; $j++){
            # Find a complimentary number to the target number
            $compNum = $target - $nums[$i]
            if ($compNum -eq $nums[$j]){
                return @($i, $j)
            }
        }
    }
}

$sum = Get-twoSums $nums $target

write-host $sum
