# cerner_2tothe5th_2021
# Variables
nums = [2, 7, 11, 18]
target = 18

def twoSums(nums, target):
    # Iterate through  nums list 
    for i in range(len(nums)):
        # Iterate through nums list + 1
        for j in range(len(nums)):
            j = i + 1
            # Find complementary number to the target
            compNum = target - nums[i]
            if compNum == nums[j]:
                return (i,j)

print (twoSums(nums, target))
