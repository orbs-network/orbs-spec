# Contracts Governance Parameters Summary 

The PoS model parameters are managed by the functionalManager

| Contract                | Parameter                                   | Default [units]                                      |
| ----------------------- | ------------------------------------------- | ---------------------------------------------------- |
| committee               | maxCommitteeSize                            | 22 [#]                                               |
| elections               | minSelfStakePercentMille                    | 0 [percent-mille]                                 |
| elections               | voteOutPercentMilleThreshold                | 70000  [percent-mille]                               |
| elections               | voteUnreadyPercentMilleThreshold            | 70000  [percent-mille]                               |
| feesAndBootstrapRewards | generalCommitteeAnnualBootstrap             | 0 [10<sup>-18</sup> DAI]                             |
| feesAndBootstrapRewards | certifiedCommitteeAnnualBootstrap           | 3000 x 10<sup>18</sup> [10<sup>-18</sup> DAI]        |
| stakingRewards          | defaultDelegatorsStakingRewardsPercentMille | 66,667  [percent-mille]                              |
| stakingRewards          | maxDelegatorsStakingRewardsPercentMille     | 66,667  [percent-mille]                              |
| stakingRewards          | annualRateInPercentMille                    | 10,000  [percent-mille]                              |
| stakingRewards          | annualCap                                   | 2^256-1 [10<sup>-18</sup> ORBS] |
| subscriptions           | genesisRefTimeDelay                         | 10,800 [sec]                                         |
| subscriptions           | minimumInitialVcPayment                     | 2^256-1 [10<sup>-18</sup> ORBS]    |

> virtua chain related paramters are not supported in polygon