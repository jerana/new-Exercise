# Definition for a undirected graph node
# class UndirectedGraphNode(object):
#     def __init__(self, x):
#         self.label = x
#         self.neighbors = []

class Solution(object):
        # @param node, a undirected graph node
        # @return a undirected graph node
        
        def dfs(self,node,hMap):
            if node in hMap:
                return node
            else:
                cloneGraph = UndirectedGraphNode(node)
                hMap[node] = cloneGraph
                for nei in node.neighbors:
                    cloneGraph.neighbors.add(self.dfs(nei,hMap))

                return cloneGraph

        def clone_graph(self,node):
            if node == NONE:
                return node
            hMap = {}
            return self.dfs(node,hMap)


