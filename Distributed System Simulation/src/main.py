from utils.logger import Logger
from load_balancer import LoadBalancer
import grpc
import simulator_pb2
import simulator_pb2_grpc

class Simulator:
    def __init__(self):
        self.logger = Logger()
        self.load_balancer = LoadBalancer()

    def start(self):
        self.logger.log("Starting Distributed System Simulation...")
        # Simulate tasks
        tasks = [f"Task {i}" for i in range(10)]
        self.load_balancer.distribute_tasks(tasks)

if __name__ == "__main__":
    simulator = Simulator()
    simulator.start()
