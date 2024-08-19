import unittest
from src.main import Simulator

class TestSimulation(unittest.TestCase):
    def test_simulation_start(self):
        simulator = Simulator()
        self.assertIsInstance(simulator, Simulator)

if __name__ == '__main__':
    unittest.main()
