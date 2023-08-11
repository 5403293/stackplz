package module

const (
    PROBE_TYPE_UPROBE     = "uprobe"
    PROBE_TYPE_KPROBE     = "kprobe"
    PROBE_TYPE_TRACEPOINT = "tracepoint"
    PROBE_TYPE_PERF       = "perf"
    PROBE_TYPE_BREAKPOINT = "breakpoint"
)

const (
    MODULE_NAME_PERF    = "PerfMod"
    MODULE_NAME_BRK     = "BrkMod"
    MODULE_NAME_STACK   = "StackMod"
    MODULE_NAME_SYSCALL = "SyscallMod"
)

// http://aospxref.com/android-11.0.0_r21/xref/bionic/libc/kernel/uapi/asm-arm/asm/perf_regs.h
const (
    PERF_REG_ARM_R0 uint32 = iota
    PERF_REG_ARM_R1
    PERF_REG_ARM_R2
    PERF_REG_ARM_R3
    PERF_REG_ARM_R4
    PERF_REG_ARM_R5
    PERF_REG_ARM_R6
    PERF_REG_ARM_R7
    PERF_REG_ARM_R8
    PERF_REG_ARM_R9
    PERF_REG_ARM_R10
    PERF_REG_ARM_FP
    PERF_REG_ARM_IP
    PERF_REG_ARM_SP
    PERF_REG_ARM_LR
    PERF_REG_ARM_PC
    PERF_REG_ARM_MAX
)

// http://aospxref.com/android-11.0.0_r21/xref/bionic/libc/kernel/uapi/asm-arm64/asm/perf_regs.h
const (
    PERF_REG_ARM64_X0 uint32 = iota
    PERF_REG_ARM64_X1
    PERF_REG_ARM64_X2
    PERF_REG_ARM64_X3
    PERF_REG_ARM64_X4
    PERF_REG_ARM64_X5
    PERF_REG_ARM64_X6
    PERF_REG_ARM64_X7
    PERF_REG_ARM64_X8
    PERF_REG_ARM64_X9
    PERF_REG_ARM64_X10
    PERF_REG_ARM64_X11
    PERF_REG_ARM64_X12
    PERF_REG_ARM64_X13
    PERF_REG_ARM64_X14
    PERF_REG_ARM64_X15
    PERF_REG_ARM64_X16
    PERF_REG_ARM64_X17
    PERF_REG_ARM64_X18
    PERF_REG_ARM64_X19
    PERF_REG_ARM64_X20
    PERF_REG_ARM64_X21
    PERF_REG_ARM64_X22
    PERF_REG_ARM64_X23
    PERF_REG_ARM64_X24
    PERF_REG_ARM64_X25
    PERF_REG_ARM64_X26
    PERF_REG_ARM64_X27
    PERF_REG_ARM64_X28
    PERF_REG_ARM64_X29
    PERF_REG_ARM64_LR
    PERF_REG_ARM64_SP
    PERF_REG_ARM64_PC
    PERF_REG_ARM64_MAX
)
