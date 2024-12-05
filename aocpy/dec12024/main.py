from collections import defaultdict


left = []
right = []
with open("input.txt", "r") as file:
    for line in file:
        parts = line.strip().split(" ")
        cleaned = [int(part) for part in parts if part]
        left.append(cleaned[0])
        right.append(cleaned[1])

# NlogN time, for sort. O(1) space
def calc_total_distance(left, right):    
    # sort in place, NlogN operation
    left.sort()
    right.sort()

    # calculate diffs
    total_diff = 0
    for x, y in zip(left, right):
        total_diff += abs(x - y)
    return total_diff

# O(N) space for dict, O(M) time for all nums in left, map is uniques only
def calc_similarity_score(left, right):
    m = defaultdict(int)
    for num in right:
        m[num] += 1
    
    similarity_score = 0
    for num in left:
        similarity_score += (num * m[num])
    return similarity_score

print(f"total distance: {calc_total_distance(left, right)}")
print(f"similarity score: {calc_similarity_score(left ,right)}")

# Total Distance: 1938424
# Similarity Score: 22014209