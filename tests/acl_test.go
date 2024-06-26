package tests

import (
	"github.com/jsiebens/ionscale/pkg/client/ionscale"
	"github.com/jsiebens/ionscale/pkg/defaults"
	"github.com/jsiebens/ionscale/tests/sc"
	"github.com/jsiebens/ionscale/tests/tsn"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestACL_PeersShouldBeRemovedWhenNoMatchingACLRuleIsAvailable(t *testing.T) {
	sc.Run(t, func(s *sc.Scenario) {
		tailnet := s.CreateTailnet()
		clientKey := s.CreateAuthKey(tailnet.Id, true, "tag:client")
		serverKey := s.CreateAuthKey(tailnet.Id, true, "tag:server")

		client1 := s.NewTailscaleNode()
		client2 := s.NewTailscaleNode()
		server := s.NewTailscaleNode()

		require.NoError(t, client1.Up(clientKey))
		require.NoError(t, client2.Up(clientKey))
		require.NoError(t, server.Up(serverKey))
		require.NoError(t, server.WaitFor(tsn.PeerCount(2)))

		policy := defaults.DefaultACLPolicy()
		policy.ACLs = []ionscale.ACLEntry{
			{
				Action:      "accept",
				Source:      []string{"tag:server"},
				Destination: []string{"tag:server:*"},
			},
		}

		s.SetACLPolicy(tailnet.Id, policy)

		require.NoError(t, server.WaitFor(tsn.PeerCount(0)))
	})
}

func TestACL_PeersShouldSeeSSHRecorder(t *testing.T) {
	sc.Run(t, func(s *sc.Scenario) {
		tailnet := s.CreateTailnet()
		clientKey := s.CreateAuthKey(tailnet.Id, true, "tag:client")
		recorderKey := s.CreateAuthKey(tailnet.Id, true, "tag:recorder")

		policy := defaults.DefaultACLPolicy()
		policy.ACLs = []ionscale.ACLEntry{
			{
				Action:      "accept",
				Source:      []string{"tag:client"},
				Destination: []string{"tag:client:*"},
			},
		}
		policy.SSH = []ionscale.ACLSSH{
			{
				Action:      "check",
				Source:      []string{"tag:client"},
				Destination: []string{"tag:client"},
				Users:       []string{"autogroup:nonroot", "root"},
				Recorder:    []string{"tag:recorder"},
			},
		}

		s.SetACLPolicy(tailnet.Id, policy)

		client1 := s.NewTailscaleNode()
		client2 := s.NewTailscaleNode()
		recorder1 := s.NewTailscaleNode()

		require.NoError(t, client1.Up(clientKey))
		require.NoError(t, client2.Up(clientKey))
		require.NoError(t, recorder1.Up(recorderKey))

		require.NoError(t, recorder1.WaitFor(tsn.PeerCount(2)))
	})
}
