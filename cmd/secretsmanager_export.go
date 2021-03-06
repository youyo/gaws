// Copyright © 2019 youyo <1003ni2@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/spf13/cobra"
)

var secretsmanagerExportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export secure string",
	Run: func(cmd *cobra.Command, args []string) {
		sess := session.Must(
			session.NewSessionWithOptions(
				session.Options{
					SharedConfigState: session.SharedConfigEnable,
				},
			),
		)
		svc := secretsmanager.New(sess)
		input := &secretsmanager.GetSecretValueInput{
			SecretId: aws.String(secretsmanagerCmdFlagsId),
		}
		result, err := svc.GetSecretValue(input)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(*result.SecretString)
	},
}

func init() {
	secretsmanagerCmd.AddCommand(secretsmanagerExportCmd)
}
