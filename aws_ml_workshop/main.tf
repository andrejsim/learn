resource "aws_s3_bucket" "examplebucket" {
  bucket = "${var.s3}"
  acl    = "private"
}

resource "aws_iam_role" "role" {
  name = "test-role"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "ec2.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF
}

resource "aws_iam_policy" "policy" {
  name        = "service-policy"
  description = "A sagemaker policy"

  policy = <<EOF
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Sid": "workshopPerm1",
            "Effect": "Allow",
            "Action": [
                "personalize:*",
                "s3:*",
                "lex:*"
            ],
            "Resource": "*"
        },
        {
            "Sid": "workshopPerm2",
            "Effect": "Allow",
            "Action": [
                "iam:GetRole",
                "iam:PassRole",
                "iam:CreateRole",
                "iam:AttachRolePolicy"
            ],
            "Resource": "arn:aws:iam::*:role/*"
        }
    ]
}
EOF
}

resource "aws_iam_role_policy_attachment" "test-attach" {
  role       = "${aws_iam_role.role.name}"
  policy_arn = "${aws_iam_policy.policy.arn}"
}

resource "aws_sagemaker_notebook_instance" "ni" {
  name          = "${var.sage}"
  role_arn      = "${aws_iam_role.role.arn}"
  instance_type = "ml.m5.xlarge"

  tags = {
    team = "balkan"
    provision = "terraform"
  }
}