// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package utils

import (
	"bufio"
	"fmt"
	"io"
)

func ScannerPrint(scanner *bufio.Scanner, writer io.Writer) error {
	if writer == nil {
		return nil
	}
	for scanner.Scan() {
		line := scanner.Text()
		if _, err := fmt.Fprintln(writer, line); err != nil {
			return err
		}
	}
	return nil
}