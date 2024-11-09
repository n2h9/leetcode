using namespace ::std;

#include <queue>
#include <unordered_map>
#include <unordered_set>
#include <utility>
#include <vector>

class Solution {
 public:
  vector<int> countPairsOfConnectableServers(vector<vector<int>>& edges,
                                             int signalSpeed) {
    // amap[x][y] = amap[y][x] = distance of edge between x and y
    auto create_adjacency_map_f = [](vector<vector<int>>& edges) -> auto {
      auto map = unordered_map<int, unordered_map<int, int>>{};

      for (auto& edge : edges) {
        map[edge[0]][edge[1]] = edge[2];
        map[edge[1]][edge[0]] = edge[2];
      }

      return move(map);
    };

    auto amap = create_adjacency_map_f(edges);
    auto n = amap.size();

    // visited_incoming_routes[x][y] = true if route from y to x processed
    // ( that is why they are incoming :) )
    // incoming routes are more convinient than (from, to) direction
    // as same set could be used to track number of processed hands (incoming
    // routes);
    // also
    // visited_incoming_routes[x][y] != visited_incoming_routes[y][x]
    auto visited_incoming_routes = unordered_map<int, unordered_set<int>>{};

    // same as visited_incoming_routes but indicates the route is in the queue
    // but this time first index is vertex from and set element is a vertex to
    auto routes_in_the_queue = unordered_map<int, unordered_set<int>>{};

    // num_of_servers_with_distance_map[vertex][incoming_vertex][distance] =
    // num_of_servers
    auto num_of_servers_with_distance_map =
        unordered_map<int, unordered_map<int, unordered_map<int, int>>>{};

    // q item is a pair of <vertex1, vertex2>, which is a route from vertex1 to
    // vertex2 similar to visited_routes, a route from vertex1 to vertex2 is not
    // the same as a route from vertex2 to vertex1
    auto q = queue<pair<int, int>>{};

    // fill the quee with the vertices with only one neighbour
    for (auto& vertex : amap) {
      if (vertex.second.size() == 1) {
        for (auto& vertex_destination : vertex.second) {
          // foreach is ok as map size is 1
          q.push({vertex.first, vertex_destination.first});
          routes_in_the_queue[vertex.first].insert(vertex_destination.first);
        }
      }
    }

    while (!q.empty()) {
      auto item = q.front();
      auto vertex_from = item.first;
      auto vertex_to = item.second;

      q.pop();
      routes_in_the_queue[vertex_from].erase(vertex_to);

      auto num_of_hands_of_vertex_from = amap[vertex_from].size();
      auto num_of_visited_hands_of_vertex_from =
          visited_incoming_routes[vertex_from].size();

      if (num_of_visited_hands_of_vertex_from <
          num_of_hands_of_vertex_from - 1) {
        // means that there are more than one unvisited hand
        // so we can not yet process this route
        q.push(item);
        routes_in_the_queue[item.first].insert(item.second);
        continue;
      }
      if ((num_of_visited_hands_of_vertex_from ==
           num_of_hands_of_vertex_from - 1) &&
          visited_incoming_routes[vertex_from].count(vertex_to) > 0) {
        // means that there is still one more unvisited hand
        // so we can not yet process this route either
        q.push(item);
        routes_in_the_queue[item.first].insert(item.second);
        continue;
      }

      // if we here then:
      // - or all incoming routes were processed;
      // - or there is one incoming route to process,
      //   but this route is opposite to current route, and does not affect the
      //   result;

      auto distance_between_from_and_to = amap[vertex_from][vertex_to];

      // propagate information on number of vertices per weight
      // from vertex_from to vertex_to
      for (auto& entity_a : num_of_servers_with_distance_map[vertex_from]) {
        auto source_vertex = entity_a.first;
        // do not put in vertex_to the values that came from vertex_to itself :)
        if (source_vertex == vertex_to) {
          continue;
        }
        auto distance_to_number_map = entity_a.second;
        for (auto& entity_b : distance_to_number_map) {
          auto distance = entity_b.first + distance_between_from_and_to;
          auto number_of_servers = entity_b.second;
          num_of_servers_with_distance_map[vertex_to][vertex_from][distance] +=
              number_of_servers;
        }
      }
      // also include current server
      num_of_servers_with_distance_map[vertex_to][vertex_from]
                                      [distance_between_from_and_to]++;

      // mark route as visited
      visited_incoming_routes[vertex_to].insert(vertex_from);

      // push unvisited neighbour routes to queue
      for (auto& neighbour : amap[vertex_to]) {
        auto neighbour_vertex = neighbour.first;
        if ((visited_incoming_routes[neighbour_vertex].count(vertex_to) <= 0) &&
            ((routes_in_the_queue[vertex_to].count(neighbour_vertex) <= 0))) {
          q.push({vertex_to, neighbour_vertex});
          routes_in_the_queue[vertex_to].insert(neighbour_vertex);
        }
      }
    }

    // filter out distances not divisable by signalSpeed
    auto num_of_servers_connected_to_vertex_divisible_by_signal_speed =
        unordered_map<int, int>{};
    auto num_of_servers_connected_to_vertex_divisible_by_signal_speed_by_hand =
        unordered_map<int, unordered_map<int, int>>{};

    for (auto& per_vertex : num_of_servers_with_distance_map) {
      auto vertex = per_vertex.first;
      auto hands = per_vertex.second;
      for (auto& hand : hands) {
        auto incoming_vertex = hand.first;
        auto distance_to_number_map = hand.second;
        for (auto& entity : distance_to_number_map) {
          auto distance = entity.first;
          auto number_of_servers = entity.second;
          if (distance % signalSpeed == 0) {
            num_of_servers_connected_to_vertex_divisible_by_signal_speed
                [vertex] += number_of_servers;
            num_of_servers_connected_to_vertex_divisible_by_signal_speed_by_hand
                [vertex][incoming_vertex] += number_of_servers;
          }
        }
      }
    }

    auto result = vector<int>();
    for (decltype(n) vertex_i = 0; vertex_i < n; vertex_i++) {
      auto hands =
          num_of_servers_connected_to_vertex_divisible_by_signal_speed_by_hand
              [vertex_i];
      auto num_of_hands = hands.size();

      // if there is only one hand -> then number of connected servers is 0
      auto iteration_res = 0;
      if (num_of_hands > 1) {
        auto total_num_of_servers =
            num_of_servers_connected_to_vertex_divisible_by_signal_speed
                [vertex_i];
        for (auto& hand : hands) {
          auto hand_num_of_servers = hand.second;
          // so this servers connects all other connected
          // to current processing hand
          iteration_res += (total_num_of_servers - hand_num_of_servers) *
                           hand_num_of_servers;
        }

        // as we iterate over all elements
        iteration_res /= 2;
      }

      result.push_back(iteration_res);
    }

    return result;
  }
};