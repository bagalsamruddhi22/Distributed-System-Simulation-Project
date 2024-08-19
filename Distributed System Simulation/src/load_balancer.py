class LoadBalancer:
    def __init__(self):
        self.nodes = ["Node 1", "Node 2", "Node 3"]

    def distribute_tasks(self, tasks):
        for i, task in enumerate(tasks):
            node = self.nodes[i % len(self.nodes)]
            self.assign_task(task, node)

    def assign_task(self, task, node):
        print(f"Assigning {task} to {node}")
