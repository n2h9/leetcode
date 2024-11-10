1) For every server (vertex) for every hand (edge) calculate the number of connected servers to it by distance: num_of_servers_with_distance_map; 

Complexity: O(n**2);

2) Transform num_of_servers_with_distance_map to:
- num_of_servers_connected_to_vertex_divisible_by_signal_speed:
- num_of_servers_connected_to_vertex_divisible_by_signal_speed_by_hand
This provides information of connected servers per vertex, and connected servers per 
vertex per hand. Also this is already only number of servers which satisfy condition on signal speed;

Complexity: O(n); 

3) Using two maps from 2) for every hand of every vertex the number of servers connected through current hand to all other hands of this vertex is 
num = (total_num_of_servers - hand_num_of_servers) * hand_num_of_servers;  

Complexity: O(n);
