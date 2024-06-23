local function solveNQueens(n)
    local solutions = {}
    local function solve(queens, row, cols, diags1, diags2)
        if row == n then
            table.insert(solutions, queens)
            return
        end
        local availablePositions = ((1 << n) - 1) & ~(cols | diags1 | diags2)
        while availablePositions > 0 do
            local position = availablePositions & -availablePositions
            availablePositions = availablePositions - position
            solve(queens .. string.char(97 + math.floor(math.log(position) / math.log(2))), row + 1, cols | position, (diags1 | position) << 1, (diags2 | position) >> 1)
        end
    end
    solve("", 0, 0, 0, 0)
    return solutions
end

-- Prompt for board size
io.write("Enter the board size (N): ")
local n = tonumber(io.read("*l"))

-- Measure the time taken to solve
local startTime = os.clock()
local solutions = solveNQueens(n)
local endTime = os.clock()

-- Print the solutions
print("Number of solutions: " .. #solutions)
for i, solution in ipairs(solutions) do
    print("Solution " .. i .. ": " .. solution)
end

-- Print the time taken
print(string.format("Time taken: %.4f seconds", endTime - startTime))
