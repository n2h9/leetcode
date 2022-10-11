# Solution

The division by modulo is associative operation
The range of k is small [1,50]
Need to calculate an array C, of number of paths per reminder, of len k for each cell.
    [i] -> v where "i" is a reminder, v - is a the number of pathes with the reminder i to this cell.

The direction of movement is right, or down.
It means that reach each cell is possbile only from left or up.

Start from the top row
Itearte over all matrix from left to right.

Only need to store the previous row array of number paths per reminder and current. A result of this row depends on element on top and left from cell.

The result is in the last element of array of number of paths per reminder with the index 0.
0 -> means reminder is 0 -> hence pathes are divided by k;

Space complexity:
O(n * k) - need to store only two rows, k array per each element per row

Time complexity:
O(n * m * k) - algorithm goes ktimes through each cell of a matrix, 

