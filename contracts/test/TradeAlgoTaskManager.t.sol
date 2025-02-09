// // SPDX-License-Identifier: UNLICENSED
// pragma solidity ^0.8.9;

// import "forge-std/Test.sol";
// import "../src/ITradeAlgoTaskManager.sol";
// import "../src/TradeAlgoTaskManager.sol";
// import "@eigenlayer-middleware/src/libraries/BN254.sol";
// import "@openzeppelin/contracts/access/Ownable.sol";

// contract TradeAlgoTaskManagerTest is Test {
//     TradeAlgoTaskManager public tradeAlgoTaskManager;
//     address public deployer;
//     address public investor;
//     address public aggregator;
//     address public generator;

//     uint256 public strategyId = 1;
//     uint32 public quorumThresholdPercentage = 100;
//     bytes public quorumNumbers;

//     function setUp() public {
//         deployer = address(this);
//         investor = address(0x123);
//         aggregator = address(0x456);
//         generator = address(0x789);

//         tradeAlgoTaskManager = new TradeAlgoTaskManager(
//             IRegistryCoordinator(address(0)), // Mocked Registry Coordinator
//             30 // Task Response Window Blocks
//         );

//         quorumNumbers = abi.encodePacked(uint8(1), uint8(2));
//     }

//     function test_TaskManager_InitializedCorrectly() public {
//         assertEq(
//             tradeAlgoTaskManager.getTaskResponseWindowBlock(),
//             30,
//             "TaskManager initialization incorrect"
//         );
//     }

//     function test_CreateNewTask() public {
//         tradeAlgoTaskManager.createNewTask(
//             strategyId, investor, quorumThresholdPercentage, quorumNumbers
//         );
//         uint32 latestTaskIndex = tradeAlgoTaskManager.taskNumber();
//         assertEq(latestTaskIndex, 1, "Task creation failed");
//     }

//     function test_CreateMultipleTasks() public {
//         tradeAlgoTaskManager.createNewTask(
//             strategyId, investor, quorumThresholdPercentage, quorumNumbers
//         );
//         tradeAlgoTaskManager.createNewTask(
//             strategyId, investor, quorumThresholdPercentage, quorumNumbers
//         );
//         tradeAlgoTaskManager.createNewTask(
//             strategyId, investor, quorumThresholdPercentage, quorumNumbers
//         );
//         uint32 latestTaskIndex = tradeAlgoTaskManager.taskNumber();
//         assertEq(latestTaskIndex, 3, "Multiple task creation failed");
//     }

//     function test_SubmitTaskResponse() public {
//         tradeAlgoTaskManager.createNewTask(
//             strategyId, investor, quorumThresholdPercentage, quorumNumbers
//         );
//         uint32 latestTaskIndex = tradeAlgoTaskManager.taskNumber();
//         assertEq(latestTaskIndex, 1, "Task creation failed");

//         ITradeAlgoTaskManager.Task memory task = ITradeAlgoTaskManager.Task({
//             strategyId: strategyId,
//             investor: investor,
//             taskCreatedBlock: uint32(block.number),
//             quorumNumbers: quorumNumbers,
//             quorumThresholdPercentage: quorumThresholdPercentage
//         });

//         ITradeAlgoTaskManager.TaskResponse memory taskResponse = ITradeAlgoTaskManager.TaskResponse({
//             referenceTaskIndex: latestTaskIndex,
//             calculatedResult: 42, // Example result
//             operatorSignature: bytes("mock_signature")
//         });

//         ITradeAlgoTaskManager.TaskResponseMetadata memory taskResponseMetadata =
//         ITradeAlgoTaskManager.TaskResponseMetadata({
//             taskRespondedBlock: uint32(block.number),
//             hashOfNonSigners: keccak256(abi.encodePacked(aggregator))
//         });

//         BN254.G1Point[] m = tradeAlgoTaskManager.submitTaskResponse(
//             task,
//             taskResponse,
//             taskResponseMetadata,
//             quorumThresholdPercentage,
//             nonSignerStakesAndSignature
//         );
//     }

//     function test_ChallengeTaskResponse() public {
//         tradeAlgoTaskManager.createNewTask(
//             strategyId, investor, quorumThresholdPercentage, quorumNumbers
//         );
//         uint32 latestTaskIndex = tradeAlgoTaskManager.taskNumber();
//         assertEq(latestTaskIndex, 1, "Task creation failed");

//         ITradeAlgoTaskManager.Task memory task = ITradeAlgoTaskManager.Task({
//             strategyId: strategyId,
//             investor: investor,
//             taskCreatedBlock: uint32(block.number),
//             quorumNumbers: quorumNumbers,
//             quorumThresholdPercentage: quorumThresholdPercentage
//         });

//         ITradeAlgoTaskManager.TaskResponse memory taskResponse = ITradeAlgoTaskManager.TaskResponse({
//             referenceTaskIndex: latestTaskIndex,
//             calculatedResult: 42,
//             operatorSignature: bytes("mock_signature")
//         });

//         ITradeAlgoTaskManager.TaskResponseMetadata memory taskResponseMetadata =
//         ITradeAlgoTaskManager.TaskResponseMetadata({
//             taskRespondedBlock: uint32(block.number),
//             hashOfNonSigners: keccak256(abi.encodePacked(aggregator))
//         });

//         BN254.G1Point[] m = tradeAlgoTaskManager.raiseAndResolveChallenge(
//             task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators
//         );
//     }
// }
