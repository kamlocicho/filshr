import { Flex, ProgressCircle } from "@adobe/react-spectrum";

export default function Loading() {
    return (
        <Flex width={"100%"} height={"100%"} alignItems={"center"} justifyContent={"center"}>
            <ProgressCircle aria-label="Loading…" isIndeterminate size="L" />
        </Flex>
    )
}
