package tlpstress

import (
	"fmt"
	"github.com/thelastpickle/stress-operator/pkg/apis/thelastpickle/v1alpha1"
	"strconv"
	"strings"
)

type argName string

const (
	run                argName = "run"
	host               argName = "--host"
	consistencyLevel   argName = "--cl"
	partitions         argName = "--partitions"
	duration           argName = "--duration"
	drop               argName = "--drop"
	iterations         argName = "--iterations"
	readRate           argName = "--readrate"
	populate           argName = "--populate"
	concurrency        argName = "--concurrency"
	partitionGenerator argName = "--partitiongenerator"
	dataCenter         argName = "--dc"
	replication        argName = "--replication"
	username           argName = "--username"
	password           argName = "--password"
	keyspace           argName = "--keyspace"
	ttl                argName = "--ttl"
	paging             argName = "--paging"
	deleteRate         argName = "--deleterate"
)

type CommandLineArgs struct {
	args []string
}

func (c *CommandLineArgs) addArg(name argName, val string) {
	c.args = append(c.args, string(name))
	c.args = append(c.args, val)
}

func (c *CommandLineArgs) GetArgs() *[]string {
	return &c.args
}

func (c *CommandLineArgs) String() string {
	return fmt.Sprint(c.args)
}

// Generates the arguments that are passed to the tlp-stress executable
func CreateCommandLineArgs(stressCfg *v1alpha1.StressConfig, cassandraCfg *v1alpha1.CassandraConfig, namespace string) *CommandLineArgs {
	args:= CommandLineArgs{}

	args.addArg(run, string(stressCfg.Workload))

	if len(stressCfg.ConsistencyLevel) > 0 {
		args.addArg(consistencyLevel, string(stressCfg.ConsistencyLevel))
	}

	if stressCfg.Partitions != nil {
		args.addArg(partitions, *stressCfg.Partitions)
	}

	if len(stressCfg.Duration) > 0 {
		args.addArg(duration, stressCfg.Duration)
	}

	if stressCfg.DropKeyspace {
		args.addArg(drop, "")
	}

	if stressCfg.Iterations != nil {
		args.addArg(iterations, *stressCfg.Iterations)
	}

	if len(stressCfg.ReadRate) > 0 {
		args.addArg(readRate, stressCfg.ReadRate)
	}

	if stressCfg.Populate != nil {
		args.addArg(populate, *stressCfg.Populate)
	}

	if stressCfg.Concurrency != nil && *stressCfg.Concurrency != 100 {
		args.addArg(concurrency, strconv.FormatInt(int64(*stressCfg.Concurrency), 10))
	}

	if len(stressCfg.PartitionGenerator) > 0 {
		args.addArg(partitionGenerator, stressCfg.PartitionGenerator)
	}

	if len(stressCfg.DataCenter) > 0 {
		args.addArg(dataCenter, stressCfg.DataCenter)
	}

	if len(stressCfg.Username) > 0 {
		args.addArg(username, stressCfg.Username)
	}

	if len(stressCfg.Password) > 0 {
		args.addArg(password, stressCfg.Password)
	}

	if len(stressCfg.Keyspace) > 0 {
		args.addArg(keyspace, stressCfg.Keyspace)
	}

	if stressCfg.Ttl != nil {
		args.addArg(ttl, strconv.FormatInt(*stressCfg.Ttl, 10))
	}

	if stressCfg.Paging != nil {
		args.addArg(paging, strconv.FormatInt(int64(*stressCfg.Paging), 10))
	}

	if len(stressCfg.DeleteRate) > 0 {
		args.addArg(deleteRate, stressCfg.DeleteRate)
	}

	// TODO Need to make sure only one replication strategy is specified
	if stressCfg.Replication.SimpleStrategy != nil {
		replicationFactor := strconv.FormatInt(int64(*stressCfg.Replication.SimpleStrategy), 10)
		replicationString := fmt.Sprintf(`{'class': 'SimpleStrategy', 'replication_factor': %s}`, replicationFactor)
		args.addArg(replication, replicationString)
	} else if stressCfg.Replication.NetworkTopologyStrategy != nil {
		var sb strings.Builder
		dcs := make([]string, 0)
		for k, v := range *stressCfg.Replication.NetworkTopologyStrategy {
			sb.WriteString("'")
			sb.WriteString(k)
			sb.WriteString("': ")
			sb.WriteString(strconv.FormatInt(int64(v), 10))
			dcs = append(dcs, sb.String())
			sb.Reset()
		}
		replicationString := fmt.Sprintf("{'class': 'NetworkTopologyStrategy', %s}", strings.Join(dcs, ", "))
		args.addArg(replication, replicationString)
	}

	// TODO add validation check that either CassandraService or CassandraCluster is defined in the spec
	svc := ""
	if cassandraCfg.CassandraClusterTemplate != nil {
		if cassandraCfg.CassandraClusterTemplate.Namespace == namespace {
			svc = cassandraCfg.CassandraClusterTemplate.Name
		} else {
			svc = fmt.Sprintf("%s.%s.svc.cluster.local", cassandraCfg.CassandraClusterTemplate.Name,
				cassandraCfg.CassandraClusterTemplate.Namespace)
		}
	} else if cassandraCfg.CassandraCluster != nil {
		// The headless service for a CassandraCluster has the same name as the cluster
		if cassandraCfg.CassandraCluster.Namespace == "" || cassandraCfg.CassandraCluster.Namespace == namespace {
			svc = cassandraCfg.CassandraCluster.Name
		} else {
			// CassandraCluster service is in a different namespace
			svc = fmt.Sprintf("%s.%s.svc.cluster.local", cassandraCfg.CassandraCluster.Name,
				cassandraCfg.CassandraCluster.Namespace)
		}
	} else {
		svc = cassandraCfg.CassandraService
	}
	args.addArg(host, svc)

	return &args
}