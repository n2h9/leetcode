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
    // amap[x][y] = amap[y][x] = weight of edge between x and y
    auto create_adjacency_map_f = [](vector<vector<int>>& edges) -> auto {
      auto map = unordered_map<int, unordered_map<int, int>>{};

      for (auto& edge : edges) {
        map[edge[0]][edge[1]] = edge[2];
        map[edge[1]][edge[0]] = edge[2];
      }

      return move(map);
    };

    // map[weight] = number of servers
    auto create_map_of_weighted_pathes_f =
        [signalSpeed](
            unordered_map<int, unordered_map<int, int>>& adjacency_map,
            int main_vertex, int neighbor_vertex) -> auto {
      auto map = unordered_map<int, int>{};

      auto weight = adjacency_map[main_vertex][neighbor_vertex];

      // queue item here is a pair of <vertex, weight from main_vertex>
      auto q = queue<pair<int, int>>{};
      auto visited = unordered_set<int>{};

      // do not return to main_vertex
      visited.insert(main_vertex);
      q.push({neighbor_vertex, weight});

      while (!q.empty()) {
        auto item = q.front();
        q.pop();

        auto v = item.first;
        auto cumulative_weight = item.second;
        visited.insert(v);

        // put all non visited relatives to q
        for (auto& relative : adjacency_map[v]) {
          auto relative_v = relative.first;
          auto relative_weight = relative.second;

          if (visited.count(relative_v) <= 0) {
            auto weight = cumulative_weight + relative_weight;
            q.push({relative_v, weight});
          }
        }

        // main action: increment number of servers per weight
        //
        // optimisation: no sense to store weights in map if cumulative_weight
        // is not divided by signalSpeed
        if (cumulative_weight % signalSpeed == 0) {
          map[cumulative_weight]++;
        }
      }

      return move(map);
    };

    // map[vertex][][weights] = num of servers
    auto create_map_of_vector_of_map_of_weighted_pathes_f =
        [&create_map_of_weighted_pathes_f](
            unordered_map<int, unordered_map<int, int>>& adjacency_map)
        -> auto {
      auto map = unordered_map<int, vector<unordered_map<int, int>>>{};

      // number of vertices
      auto n = adjacency_map.size();

      for (decltype(n) vertex_i = 0; vertex_i < n; vertex_i++) {
        for (auto& neighbor : adjacency_map[vertex_i]) {
          map[vertex_i].push_back(create_map_of_weighted_pathes_f(
              adjacency_map, vertex_i, neighbor.first));
        }
      }

      return move(map);
    };

    auto amap = create_adjacency_map_f(edges);

    // map[vertex][hand][weights] = num_of_servers
    auto num_of_servers_with_weight_map =
        create_map_of_vector_of_map_of_weighted_pathes_f(amap);

    // number of vertex;
    auto n = amap.size();

    auto result = vector<int>();
    for (decltype(n) vertex_i = 0; vertex_i < n; vertex_i++) {
      auto hands = num_of_servers_with_weight_map[vertex_i];
      auto num_of_hands = hands.size();

      // if there is only one hand -> then number of ocnnected servers is 0
      auto iteration_res = 0;
      if (num_of_hands > 1) {
        for (decltype(num_of_hands) i = 0; i < num_of_hands; i++) {
          for (auto& entity_i : hands[i]) {
            // auto weight_i = entity_i.first;
            auto num_of_servers_i = entity_i.second;
            for (decltype(i) j = i + 1; j < num_of_hands; j++) {
              for (auto& entity_j : hands[j]) {
                // auto weight_j = entity_j.first;
                auto num_of_servers_j = entity_j.second;
                // optimisation: the check on weight divded by signalSpeed
                // does not make sense here as only weights divided by signal
                // speed are in the map
                // if ((weight_i % signalSpeed == 0) &&
                //     (weight_j % signalSpeed == 0)) {
                iteration_res += num_of_servers_i * num_of_servers_j;
                // }
              }
            }
          }
        }
      }

      result.push_back(iteration_res);
    }

    return result;
  }
};