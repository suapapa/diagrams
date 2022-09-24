from diagrams import Diagram
from diagrams.aws.compute import EC2
from diagrams.aws.database import RDS
from diagrams.aws.network import ELB
import time

# Diagrams Sandbox: DO NOT DELETE THIS LINE #

time.sleep(365*24*60*60)
with Diagram("Web Service", show=False):
        ELB("lb") >> EC2("web") >> RDS("userdb")
