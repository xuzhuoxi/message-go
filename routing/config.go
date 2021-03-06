package routing

import (
	"sync"
)

type RoutingSetting struct {
	Id      string
	Mode    RoutingMode
	Formats []string
}

//----------------------

type StrategyConfig struct {
	Targets []IRoutingTarget
	Mu      sync.RWMutex
}

func (s *StrategyConfig) TargetSize() int {
	s.Mu.RLock()
	defer s.Mu.RUnlock()
	return len(s.Targets)
}

func (s *StrategyConfig) AppendRoutingTarget(target IRoutingTarget) error {
	if nil == target {
		return ErrRoutingTargetNil
	}
	s.Mu.Lock()
	defer s.Mu.Unlock()
	s.Targets = append(s.Targets, target)
	return nil
}

func (s *StrategyConfig) AppendRoutingTargets(targets []IRoutingTarget) error {
	if len(targets) == 0 {
		return ErrRoutingTargetsEmpty
	}
	for idx, _ := range targets {
		if nil == targets[idx] {
			return ErrRoutingTargetNil
		}
	}
	s.Mu.Lock()
	defer s.Mu.Unlock()
	s.Targets = append(s.Targets, targets...)
	return nil
}

func (s *StrategyConfig) SetRoutingTargets(targets []IRoutingTarget) error {
	if len(targets) == 0 {
		return ErrRoutingTargetsEmpty
	}
	for idx, _ := range targets {
		if nil == targets[idx] {
			return ErrRoutingTargetNil
		}
	}
	s.Mu.Lock()
	defer s.Mu.Unlock()
	s.Targets = targets
	return nil
}
